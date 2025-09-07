package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value atomic.Uint64
}

func (c *Counter) Inc() {
	c.value.Add(1)
}

func (c *Counter) Load() uint64 {
	return c.value.Load()
}

func main() {
	counter := &Counter{}

	wg := &sync.WaitGroup{}

	for range 50 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Count:", counter.Load())
}
