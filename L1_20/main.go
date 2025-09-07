package main

import (
	"fmt"
	"strings"
)

func reverseSentence(s string) string {
	sentence := strings.Split(s, " ")

	for i := 0; i < len(sentence)/2; i++ {
		sentence[i], sentence[len(sentence)-1-i] = sentence[len(sentence)-1-i], sentence[i]
	}

	return strings.Join(sentence, " ")
}

func main() {
	sentence := "snow dog sun"
	fmt.Println(reverseSentence(sentence))
}
