package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func err1() {
	fmt.Println("\nerr1----------------------------------------")
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func err2() {
	fmt.Println("\nerr2----------------------------------------")
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func err3() {
	fmt.Println("\nerr3----------------------------------------")
	addr, err := net.LookupHost("golangbot123.com")

	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func err4() {
	fmt.Println("\nerr4----------------------------------------")
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func main() {
	err1()
	err2()
	err3()
	err4()
}
