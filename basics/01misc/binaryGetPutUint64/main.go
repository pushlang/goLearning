package main

import (
	"encoding/binary"
	"fmt"
	//"bytes"
)

func main() {
	a := uint64(0x00_00_00_00_00_0F_42_40)
	b := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x42, 0x40}

	//fmt.Printf("%064b\n", a)

	for i := 0; i < 8; i++ {
		a <<= 8
		a |= uint64(b[i])
	}

	fmt.Printf("%064b\t%[1]d\n", a)

	isUint64 := getUint64(b)
	fmt.Printf("%064b\t%[1]d\n\n", isUint64)

	isByte := putUint64(isUint64)
	fmt.Printf("%08b\t%[1]d\n\n", isByte)

}

func getUint64(s []byte) uint64 {
	b:=make([]byte, len(s))
	copy(b[8-len(s):], s)
	return uint64(binary.BigEndian.Uint64(b[:]))
}

func putUint64(ui uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b[0:], ui)
	return b
}
