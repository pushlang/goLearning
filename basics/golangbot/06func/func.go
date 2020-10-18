package main

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) { //parameters same type
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter //multiple return values
}

func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	area, perimeter := rectProps(10.8, 5.6) //multiple return values
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	area2, _ := rectProps2(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area2)
}
