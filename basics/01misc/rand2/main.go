package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	rand.Read(b)

	return b, nil
}

func main() {
	token, _ := GenerateRandomBytes(1)
	fmt.Printf("%d", token)
}