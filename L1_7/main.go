// go run ./L1_7/main.go -race
package main

import (
	"fmt"
	"sync"
)

type threadSafeMap struct {
	mu sync.RWMutex
	m  map[interface{}]interface{}
}

func newThreadSafeMap() *threadSafeMap {
	return &threadSafeMap{
		mu: sync.RWMutex{},
		m:  make(map[interface{}]interface{}),
	}
}

func (m *threadSafeMap) Set(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key] = value
}

func (m *threadSafeMap) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.m[key]

	return value, ok
}

func main() {
	m := newThreadSafeMap()
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			m.Set(i, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			value, ok := m.Get(i)
			if !ok {
				fmt.Println("Not found")
				continue
			}

			fmt.Println(value)
		}
	}()

	wg.Wait()
}
