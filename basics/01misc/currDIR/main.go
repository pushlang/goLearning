package main

import (
	"log"
	"os"
	"fmt"
)

func main() {

	cwd, err := os.Getwd() // Ошибка компиляции: cwd не используется
	if err != nil {
		log.Fatalf("Ошибка os.Getwd: %v", err)
	}
	fmt.Printf(cwd)

}
