package main

import (
	"fmt"
)

func reverseString(s string) string {
	inputStr := []rune(s)
	left, right := 0, len(inputStr)-1

	for left < right {
		inputStr[left], inputStr[right] = inputStr[right], inputStr[left]
		left++
		right--
	}

	return string(inputStr)
}

func main() {
	str := "Привет, Мир!"

	fmt.Println(reverseString(str))
}
