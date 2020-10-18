package main

import (
	"fmt"
)

func main() {
	s1 := [][]string{{"", "H", "He", "Hel", "Hello", "Hello World!"},
		{"", "H", "He", "Hel", "Hello", "Hello World!"},
		{"", "H", "He", "Hel", "Hello", "Hello World!"},
		{"", "H", "He", "Hel", "Hello", "Hello World!"},
		{"0", "01", "012", "0123", "01234", "012345 6789"},
	}
	s2 := [][]string{{"", "H", "He", "Hel", "Hello", "Hello World!"},
		{"", "H", "eH", "eHl", "elHlo", " Wd!Hoellorl"},
		{" ", "h", "eHl", "eH", "loelH", " Wd!Hellorl"},
		{"", "", "H", "He", "Hel", "Hello World"},
		{"0", "10", "120", "1230", "34012", "0123 456789"},
	}

	for i, v1 := range s1 {
		for j := range v1 {
			fmt.Printf("%20s\t%20s\t%t\n", s1[i][j], s2[i][j], anagram(s1[i][j], s2[i][j]))
		}
		fmt.Println("")
	}

}

func anagram(s1, s2 string) bool {
	var map1 map[rune]int
	var map2 map[rune]int

	if l1, l2 := len(s1), len(s2); l1 != l2 || l1 == 0 || l2 == 0 {
		return false
	}

	for _, v := range s1 {
		if _, ok := map1[v]; ok {
			map1[v]++
			map2[v]++
		}
	}

	for k1, v1 := range map1 {
		if v2, _ := map2[k1]; v1 != v2 {
			return false
		}
	}

	return true
}
