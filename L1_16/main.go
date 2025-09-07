package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	// Взял в среднем всё.
	left := make([]int, 0, len(arr)/2)
	middle := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)

	for i := 0; i < len(arr); i++ {
		switch {
		case arr[i] < pivot:
			left = append(left, arr[i])
		case arr[i] == pivot:
			middle = append(middle, arr[i])
		default:
			right = append(right, arr[i])
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {
	nums := []int{4, 5, 12, 1, 0, -1123}
	fmt.Printf("Sorted numbers: %v", quickSort(nums))
}
