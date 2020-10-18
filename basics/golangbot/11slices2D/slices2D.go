package main

import "fmt"

func main() {
	{ // using append
		dim := 5
		matrix := make([][]int, dim) // dim*dim matrix
		for i := 0; i < dim; i++ {
			matrix[i] = make([]int, 0, dim)
			vector := make([]int, dim)
			for j := 0; j < dim; j++ {
				vector[j] = i*dim + j
				matrix[i] = append(matrix[i], vector[j])
			}
		}
		fmt.Println(matrix)
	}
}
