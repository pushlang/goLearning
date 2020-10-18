package main

import (
	"fmt"
	"math"
)



func test() int {
	return 0
}

func main() {


	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
	fmt.Println(math.IsNaN(z), math.IsNaN(-z), math.IsNaN(1/z), math.IsNaN(-1/z), math.IsNaN(z/z))
	fmt.Println(math.IsInf(z, 1), math.IsInf(-z, -1), math.IsInf(1/z, 1), math.IsInf(-1/z, -1), math.IsInf(z/z, 1))
	fmt.Println(math.Log(-1.0), math.Log(1.0))
	fmt.Println(math.IsNaN(math.Log(-1.0)), math.IsNaN(math.Log(1.0)))

	fmt.Println(math.NaN())

	fmt.Println(func(z float64) (float64, bool) {
		if r := math.IsNaN(z / z); r {
			return math.NaN(), true
		}
		return z, false
	}(z))
}
