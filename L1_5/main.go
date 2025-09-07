package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func consumer(ctx context.Context, nums <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-nums:
			if !ok {
				return
			}
			time.Sleep(200 * time.Millisecond)

			fmt.Printf("Consumer got %d\n", num)
		}
	}
}

func producer(ctx context.Context, nums chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		case nums <- i:
		}
	}
}

func main() {
	numOfSeconds := flag.Int("second", 5, "number of seconds")
	flag.Parse()

	ctx, stop := context.WithTimeout(context.Background(), time.Duration(*numOfSeconds)*time.Second)
	defer stop()

	nums := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		consumer(ctx, nums)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(nums)

		producer(ctx, nums)
	}()

	wg.Wait()

	fmt.Println("Time's up. All goroutines finished")
}
