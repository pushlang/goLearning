package main

import (
	"fmt"
	"math/rand"
)

type Address struct {
	city string
}

func main() {
	// p := []Address{{"a"}, {"b"}}
	// p1 := &Address{"c"}
	// s := [][]byte{{'a', 'b', 'c'}, {'b', 'c', 'd'}}
	// fmt.Println(p[0].city)
	// fmt.Println(p1.city)
	// fmt.Printf("%s", s[:])

	for i := 0; i < 10; i++ {
		a := uint8(rand.Intn(3))
		fmt.Println(a)
	}

	var u uint8 = 255
	var v uint8 = 0
	
	for i:=0;i<100000000;i++{
	v++
		fmt.Printf("%d\r",v)
	}
	
	fmt.Println(u, u+1, u*u) // "255 0 1"
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

}
