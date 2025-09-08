package main

import (
	"fmt"
	"unicode"
)

func allSymbolsUnique(s string) bool {
	dict := make(map[rune]struct{})

	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := dict[r]; ok {
			return false
		}

		dict[r] = struct{}{}
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("str: %s, allSymbolsUnique: %t\n", str1, allSymbolsUnique(str1))
	fmt.Printf("str: %s, allSymbolsUnique: %t\n", str2, allSymbolsUnique(str2))
	fmt.Printf("str: %s, allSymbolsUnique: %t\n", str3, allSymbolsUnique(str3))
}
