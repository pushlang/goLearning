package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello world 1")
	}
	a()
	fmt.Printf("%T\n", a)

	func() {
		fmt.Println("hello world 2")
	}()

	func(n string) {
		fmt.Println("hello world 3:", n)
	}("Gophers")
}
