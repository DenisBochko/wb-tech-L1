package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.name)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La ", lyrics)
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{"James", 20}}

	action.SayHi()
	action.Sing("OOOOOOO...")
}
