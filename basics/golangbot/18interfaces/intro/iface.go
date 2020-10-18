//INTERFACE is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

package main

import (
	"fmt"
)

//VowelsFinder : interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

//MyString : MyString
type MyString string

//FindVowels : MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
	
	fmt.Printf("Vowels are %c", name.FindVowels())
}
