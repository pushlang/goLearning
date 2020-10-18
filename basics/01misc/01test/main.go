package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Printf("Hello  world!!\r")
	// time.Sleep(2 * time.Second)
	// fmt.Printf("Helol  world!!\r")
	// fmt.Printf("\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b")
	// time.Sleep(2 * time.Second)
	// fmt.Printf("                                      ")
	var str11 = []rune("привет") //len 6
	var str22 = "привет"         //len 12
	fmt.Printf("%d\n", len(str11))
	fmt.Printf("%d\n", len(str22))
	str11 = []rune(strings.ToUpper(string(str11)))
	fmt.Printf("%s\n", string(str11))
	str22 = strings.ToUpper(str22[0:])
	fmt.Printf("%s\n", str22)
	// var str2 = []rune("ПРИВЕТ世")
	// str3 := []rune("世界你好")

	// for i := 0; i < len(str1); i++ {
	// 	//fmt.Printf("%d\t%c\t%08b\n", i, str1[i], str1[i])
	// 	//fmt.Printf("%d\t%c\t%16b\n", i, str1[i], str1[i])
	// 	fmt.Printf("%d\t%c\t%16b\n", i, str1[i]-32, str1[i]-32)
	// }

	// fmt.Printf("%U\n", str1[:])
	// fmt.Printf("%U\n", str2[:])
	// fmt.Printf("%U, %U, %d\n", str1[0], str2[0], str1[0]-str2[0])
	// fmt.Printf("%c\n", str1[0]-32)
	// fmt.Printf("%U\n", str3[0])

	// str1 := "Привет мир!"
	// str2 := "мир"
	// str3 := "рим"

	// fmt.Printf("%t\n", Contains(str1, str2))
	// fmt.Printf("%t", Contains(str1, str3))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
