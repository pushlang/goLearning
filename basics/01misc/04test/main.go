package main

import "fmt"

func main() {
	fmt.Println(
		equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { // !ok
			return false // Accessing to not present item returns zero
		}
	}
	return true
}
