package main

import (
	"fmt"
)

func producer(nums []int, out chan<- int) {
	defer close(out)
	for _, n := range nums {
		out <- n
	}
}

func doubler(in <-chan int, out chan<- int) {
	defer close(out)
	for x := range in {
		out <- x * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	first := make(chan int)
	second := make(chan int)

	go producer(nums, first)
	go doubler(first, second)

	for v := range second {
		fmt.Println(v)
	}
}
