package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// type test struct {
	// 	s int
	// }

	// fmt.Printf("%#x\n", 127)
	// fmt.Printf("%#08b\n", 127)

	// fmt.Printf("%#v\n", test{0})
	// fmt.Printf("%v\n", test{0})
	// fmt.Printf("%+v\n", test{0})
	// fmt.Printf("%T\n", test{0})

	// fmt.Printf("%e\n", 0.000000000000000001)
	// fmt.Printf("%2.2E\n", 100000000000000000.)

	// a:=fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	// fmt.Println(a)
	// a=fmt.Sprintf("%6.2f", 12.0)
	// fmt.Println(a)

	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}

}
