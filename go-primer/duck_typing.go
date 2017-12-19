package main

import (
	"fmt"
)

type Animal struct {
}

func (a Animal) Quacks() {
	fmt.Println("The animal quacks")
}

func Quack(duck Duck) {
	duck.Quacks()
}

func main() {
	a := Animal{}
	Quack(a)
}
