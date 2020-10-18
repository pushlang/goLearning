package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := true, false
	c, d := a && b, a || b

	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)

	fmt.Println("\nboolean---------------------------------------------------")
	e, f := 1, 2
	res1, res2 := e > f, e < f
	fmt.Println(e, f, res1, res2)

	fmt.Println("\nint-------------------------------------------------------")
	var g int = 89
	h := 95
	fmt.Println(g, h)
	/* The output will vary on a 32(4 bytes)/64(8 bytes) bit OS. */
	fmt.Printf("%T, %d", g, unsafe.Sizeof(h))   //type and size of a
	fmt.Printf("\n%T, %d", g, unsafe.Sizeof(h)) //type and size of b

	fmt.Println("\nfloat-----------------------------------------------------")
	a2, b2 := 5.67, 8.67
	fmt.Printf("type of a %T b %T\n", a2, b2)

	sum, diff := a2+b2, a2-b2
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	fmt.Println("\ncomplex----------------------------------------------------")
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	fmt.Println("\nstring----------------------------------------------------")
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	fmt.Println("\ntype Conversion-------------------------------------------")
	i, j := 55, 67.8   //int, float64
	sum2 := i + int(j) //j is converted to int
	fmt.Println(sum2)

	var j2 float64 = float64(i) //will not work without explicit conversion
    fmt.Println("j", j2)
}
