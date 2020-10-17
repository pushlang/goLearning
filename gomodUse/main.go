package main

import (
    "fmt"

    //testmodML "github.com/pushlang/testmod"
	"github.com/pushlang/testmod/v2" 
)

func main() {
    //fmt.Println(testmodML.Hi("roberto"))
    g, err := testmod.Hi("Roberto", "pt")
	
    if err != nil {
        panic(err)
    }
    fmt.Println(g)
	g, err = testmod.Hi("Roberto", "fr")
	fmt.Println(g)
}