package main

import (
	"fmt"
	"reflect"
)

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("unknown")
		}
	}
}

func main() {
	printType(6)
	printType("Hi!")
	printType(true)
	printType(make(chan int))
}
