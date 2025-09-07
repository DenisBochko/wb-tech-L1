package main

import (
	"fmt"
)

func main() {
	// сложение/вычитание
	a, b := 5, 10

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a =", a, "b =", b)

	// XOR обмен (применим только к целым числам)
	a, b = 5, 10

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a =", a, "b =", b)

	// обычный человеческий способ
	a, b = 5, 10
	a, b = b, a

	fmt.Println("a =", a, "b =", b)
}
