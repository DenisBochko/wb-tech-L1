package main

import (
	"fmt"
)

func uniqueSet(initialSet []string) []string {
	set := make(map[string]bool)
	resultSet := make([]string, 0, len(initialSet))

	for _, word := range initialSet {
		_, ok := set[word]
		if !ok {
			set[word] = true
			resultSet = append(resultSet, word)
		}
	}

	return resultSet
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(uniqueSet(strs))
}
