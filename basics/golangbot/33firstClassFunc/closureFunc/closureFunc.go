package main

import (
	"fmt"
)

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	/* Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.*/
	aa := 5
	func() {
		fmt.Println("aa =", aa)
	}()

	/* Every closure is bound to its own surrounding variable.  */
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
