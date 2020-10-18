package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	distribution := make(map[string]int)
	sum:=0
	for _, name := range users {
		for _, letter := range name {
			switch letter {
			case 'A','a':
				distribution[name] += 1
			case 'E','e':
				distribution[name] += 1
			case 'I','i':
				distribution[name] += 2
			case 'O','o':
				distribution[name] += 3
			case 'U','u':
				distribution[name] += 4
			}
			if distribution[name] > 10 {
				distribution[name] = 10
			}
		}
		sum += distribution[name]
	}
	coins := 50 - sum
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
