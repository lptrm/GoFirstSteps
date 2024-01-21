package main

import (
	"fmt"
)

func gfMultiply(a, b, primitive byte) byte {
	var res byte

	for a != 0 {
		if a&1 == 1 {
			res ^= b
		}
		a >>= 1
		b <<= 1
	}
	for res > 8 {
		primitiveShifted := primitive
		for primitiveShifted<<1 < res {
			primitiveShifted <<= 1
		}
		res ^= primitiveShifted
	}
	return res
}

func main() {
	primitive := 0b1011
	fmt.Println("Multiplication table for GF(2^8) with primitive polynomial", primitive)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			result := gfMultiply(byte(i), byte(j), byte(primitive))
			fmt.Printf("%02x ", result)
		}
		fmt.Println()
	}
}
