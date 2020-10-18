package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := "123"
	// Convert string to int.
	number1, _ := strconv.ParseInt(value, 10, 0)
	// We now have an int.
	fmt.Println(number1)
	if number1 == 123 {
		fmt.Println(true)
	}

	number2, _ := strconv.Atoi(value)

	fmt.Println(number2)
	if number2 == 123 {
		fmt.Println(true)
	}
}
