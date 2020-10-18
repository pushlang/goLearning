package main

import (
	"fmt"
	"math"
)

func main() {
	var age int // var declaration
	age = 29    //assignment
	age = 54    //assignment

	var age2 int = 29 // var declaration with initial value

	var age3 = 29 // type will be inferred

	fmt.Println(age, age2, age3)

	var ( //declare variables in a single statement
		width, height   int = 100, 50 //declaring multiple vars
		width2, height2     = 100, 50 //"int" is dropped
		width3, height3 int           //0 assigned as their initial value
	)
	fmt.Println(width, width2, width3, height, height2, height3)

	name, age4 := "naveen", 29 //short hand declaration
	fmt.Println(name, age4)

	a, b := 20., 30.    // declare vars a and b
	b, c := 40, 50      // b is already declared but c is new
	b, c = 80, 90       // assign new values to already declared vars b and c
	d := math.Min(a, b) //assigned values which are computed

	fmt.Println(a, b, c, d)

	var aaa int = 1
	var _ = aaa // error silenced
}
