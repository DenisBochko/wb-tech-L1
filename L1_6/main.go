// go run main.go --demo=flag|channel|context-cancel|context-timeout|close-jobs|timer|goexit|signal
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	demo := flag.String("demo", "flag", "flag|channel|context-cancel|context-timeout|close-jobs|timer|goexit|signal")
	flag.Parse()

	switch *demo {
	case "flag":
		demoFlagAtomic()
	case "channel":
		demoDoneChannel()
	case "context-cancel":
		demoContextCancel()
	case "context-timeout":
		demoContextTimeout()
	case "close-jobs":
		demoCloseJobs()
	case "timer":
		demoTimerInside()
	case "goexit":
		demoGoexit()
	case "signal":
		demoSignalContext()
	default:
		fmt.Println("Unknown flag:", *demo)
		os.Exit(2)
	}
}

// 1) Выход по условию (флаг), безопасно — через sync/atomic.
// Без atomic или мьютекса чтение/запись bool из разных горутин была бы data race.
func demoFlagAtomic() {
	fmt.Println("demo: flag (atomic)")
	var stop atomic.Bool

	go func() {
		for {
			if stop.Load() {
				fmt.Println("stop by atomic flag")
				return
			}
			fmt.Println("work")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	stop.Store(true)
	time.Sleep(200 * time.Millisecond)
}

// 2) Остановка через done-канал.
func demoDoneChannel() {
	fmt.Println("demo: done channel")
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				// Прочитаем zero value из закрытого канала и завершим работу горутины.
				fmt.Println("stop by done channel")
				return
			default:
				fmt.Println("work")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(200 * time.Millisecond)
}

// 3) Контекст с ручной отменой.
func demoContextCancel() {
	fmt.Println("demo: context.WithCancel")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop by cancel()")
				return
			default:
				fmt.Println("work")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(200 * time.Millisecond)
}

// 4) Контекст с таймаутом.
func demoContextTimeout() {
	fmt.Println("demo: context.WithTimeout")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop by context timeout")
				return
			default:
				fmt.Println("work")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
}

// 5) Закрытие рабочего канала: воркер завершится, когда работа закончилась.
func demoCloseJobs() {
	fmt.Println("demo: close jobs channel")
	jobs := make(chan int)

	go func() {
		for j := range jobs {
			fmt.Println("got job", j)
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Println("worker done (input closed)")
	}()

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	time.Sleep(300 * time.Millisecond)
}

// 6) Самоостановка по таймеру внутри горутины.
func demoTimerInside() {
	fmt.Println("demo: time.After inside goroutine")
	stopAfter := time.After(1 * time.Second)

	go func() {
		for {
			select {
			case <-stopAfter:
				fmt.Println("stop by internal timer")
				return
			default:
				fmt.Println("work")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
}

// 7) Форсированно завершить текущую горутину: runtime.Goexit()
// Выполняет defers, но останавливают всю дальнейшую работу горутины.
func demoGoexit() {
	fmt.Println("demo: runtime.Goexit")
	done := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("defer before Goexit")
			close(done)
		}()

		fmt.Println("before Goexit()")
		runtime.Goexit()
		fmt.Println("after Goexit()")
	}()

	<-done
}

// 8) Сигналы ОС + контекст для graceful shutdown.
func demoSignalContext() {
	fmt.Println("demo: signal.NotifyContext (нажмите Ctrl+C или ждите авто-стоп)")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop by OS signal")

				return
			default:
				fmt.Println("work")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()

	time.Sleep(200 * time.Millisecond)
}
