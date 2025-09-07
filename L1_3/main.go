// Запуск: run: go run main.go --n=10
package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, nums <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-nums:
			if !ok {
				return
			}

			fmt.Printf("Worker %d got num %d\n", id, num)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	n := flag.Int("n", 3, "number of workers")
	flag.Parse()

	nums := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < *n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(ctx, i, nums)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				close(nums)
				return
			case nums <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Waiting for all goroutines to finish")

	wg.Wait()
}
