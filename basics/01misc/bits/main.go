package main

import (
	"fmt"
)

type bits uint8

// const (
// 	F0 bits = 1 << iota
// 	F1
// 	F2
// )

func set(b *bits, flag bits)    { *b |= flag }
func clear(b *bits, flag bits)  { *b &^= flag }
func toggle(b *bits, flag bits) { *b ^= flag }
func has(b, flag bits) bool     { return b&flag != 0 }

func main() {
	var b bits
	b = 128
	//b = set(b, F0)
	//b = toggle(b, F2)
	// for i := 0; i < 8; i++ {
	// 	*b ^= (1<<i))
	// }
	// for i, flag := range []bits{F0, F1, F2} {
	// }

	for i := 0; i < 8; i++ {
		toggle(&b, bits(1<<i))
	}
	
	fmt.Printf("%08b", b)

	{
		fmt.Println("\n------------------------------------")
		var x uint8 = 1<<1 | 1<<5
		var y uint8 = 1<<1 | 1<<2
		fmt.Printf("%08b\n", x)    //"00100010", множество {1,5}
		fmt.Printf("%08b\n", y)    //"00000110", множество {1,2}
		fmt.Printf("%08b\n", x&y)  //"00000010", пересечение {1}
		fmt.Printf("%08b\n", x|y)  //"00100110", объединение {1,2,5}
		fmt.Printf("%08b\n", x^y)  //"00100100", симметричная разность {2,5}
		fmt.Printf("%08b\n", x&^y) // "00100000", разность {5}

		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 { // Проверка принадлежности множеству
				fmt.Println(i) // "1", "5"
			}
		}

		fmt.Printf("%08b\n", x<<1) // "01000100", множество {2,6}
		fmt.Printf("%08b\n", x>>1) //"00010001", множество {0,4}

	}

}
