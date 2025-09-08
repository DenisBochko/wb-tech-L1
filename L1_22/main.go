package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Максимально возможные числа для int64 для примера
	num1 := big.NewInt(math.MaxInt64)
	num2 := big.NewInt(math.MaxInt64)
	result := big.NewInt(0)

	fmt.Printf("number 1: %d\nnumber 2: %d\n", num1, num2)
	fmt.Printf("Addition: %s\n", result.Add(num1, num2))
	fmt.Printf("Subtraction: %s\n", result.Sub(num1, num2))
	fmt.Printf("Multiplication: %s\n", result.Mul(num1, num2))
	fmt.Printf("Division: %s\n", result.Div(num1, num2))
}
