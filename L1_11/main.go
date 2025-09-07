package main

import (
	"fmt"
)

func intersectionOfSubsets(setA, setB []int) []int {
	set := make(map[int]bool)
	intersection := make([]int, 0, len(setA))

	for _, num := range setA {
		set[num] = true
	}

	for _, num := range setB {
		_, ok := set[num]
		if ok {
			intersection = append(intersection, num)
		}
	}

	return intersection
}

func main() {
	setA := []int{1, 2, 3}
	setB := []int{2, 3, 4}

	fmt.Printf("Set A: %v\nSet B: %v\nIntersection: %v\n", setA, setB, intersectionOfSubsets(setA, setB))
}
