package main

import (
	"context"
	"fmt"
	"sync"
)

func run(ctx context.Context, workerCount int, input <-chan int, transform func(num int) int) <-chan int {
	result := make(chan int, 10)
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case num, ok := <-input:
					if !ok {
						return
					}

					result <- transform(num)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func squaring(num int) int {
	return num * num
}

func main() {
	ctx := context.Background()

	nums := [...]int{2, 4, 6, 8, 10}
	numChan := make(chan int, 10)

	for _, num := range nums {
		numChan <- num
	}
	close(numChan)

	resChan := run(ctx, 5, numChan, squaring)

	for num := range resChan {
		fmt.Println(num)
	}
}
