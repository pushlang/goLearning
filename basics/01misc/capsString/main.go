package main

import (
	"fmt"
	"strings"
)

func main() {
	//x := "hello!"
	x := []rune("привет!")
	fmt.Printf("%s\n", strings.ToUpper(string(x)))

	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			//x := x + 'A' - 'a'
			//fmt.Printf("%c", x)
			fmt.Printf("%s\n", strings.ToUpper(string(x)))
		}
	}

}
