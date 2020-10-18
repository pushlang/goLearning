package main

import (
	"fmt"
	"runtime"
)

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	fmt.Println("MaxParallelism: ", maxParallelism())
}
