package main

import (
	"fmt"
)

func simple(a func(a, b int) int) {
	fmt.Printf("%X:%d\n", a, a(60, 7))
}

func main() {
	f1 := func(a, b int) int {
		return a + b
	}
	f2 := func(a, b int) int {
		return a * b
	}

	f := f1
	simple(f)
	f = f2
	simple(f)
}
