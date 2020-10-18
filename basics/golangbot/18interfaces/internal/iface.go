//INTERFACE can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.

package main

import (
	"fmt"
)

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() 
{
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()
}
