package main

import (
	"fmt"
)

func deleteItem(src *[]int, id int) {
	s := *src
	copy(s[id:], s[id+1:])
	*src = s[:len(s)-1]
}

func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("slice: %v, len: %d, cap: %d\n", src, len(src), cap(src))

	deleteItem(&src, 2)
	fmt.Printf("slice: %v, len: %d, cap: %d\n", src, len(src), cap(src))
}
