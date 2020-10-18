package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}
