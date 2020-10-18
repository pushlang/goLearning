package main

import (
	"fmt"
)

func main() {
	//The variables declared in a for loop are only available within the scope of the loop. The post statement will be executed after each successful iteration of the loop. After the post statement is executed, the condition will be rechecked. If it's true, the loop will continue executing, else the for loop terminates.
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	//nested for-loops
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//loop is terminated if i > 5
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop\n")

	//continue used to skip the current iteration
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	fmt.Println("")

outer: //labels
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}

	//multiple initialisation and increment
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//infinite loop
	i = 0
	for {
		i++
		fmt.Println("Hello World")
		if i == 5 {
			break
		}

	}

}
