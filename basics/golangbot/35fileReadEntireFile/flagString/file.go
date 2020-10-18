package main

import (
	"flag"
	"fmt"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}

//flagString.exe -fpath=c:\
