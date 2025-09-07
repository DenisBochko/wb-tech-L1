package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{-10, -3, 0, 1, 4, 5, 12, 19, 42, 123, 6666}
	fmt.Println(binarySearch(nums, 5))   // 5
	fmt.Println(binarySearch(nums, 19))  // 7
	fmt.Println(binarySearch(nums, 100)) // -1 (нет в массиве)
}
