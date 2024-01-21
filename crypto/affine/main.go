package main

import "fmt"

func main() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := make([]int, 0)
	for i := 0; i < len(alphabet); i++ {
		if gcd(i, 26) == 1 {
			s = append(s, i+65)
		}
	}
	for _, symbol := range s {
		fmt.Println(string(int32(symbol)), symbol-65)
	}
	plaintext := "ATTACK"
	a := 9
	b := 13
	ciphertext := ""
	for _, symbol := range plaintext {
		ciphertext += string(rune(encrypt(a, b, int(symbol-65)) + 65))
	}
	fmt.Println(ciphertext)
	decryptedText := ""
	for _, symbol := range ciphertext {
		decryptedText += string(rune(decrypt(a, b, int(symbol-65)) + 65))
	}
	fmt.Println(decryptedText)
}
func encrypt(a int, b int, x int) int {
	return (a*x + b) % 26
}
func modInverse(a int, m int) int {
	if a < 0 {
		a = (a % m) + m
	}
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return -1
}
func decrypt(a int, b int, y int) int {
	inverse := modInverse(a, 26)
	if inverse == -1 {
		return -1
	}
	res := inverse * (y - b) % 26
	if res < 0 {
		res += 26
	}
	return res
}
func gcd(a int, b int) int {
	if a == 0 {
		return abs(b)
	}
	if b == 0 {
		return abs(a)
	}
	for b != 0 {
		h := a % b
		a = b
		b = h

	}
	return abs(a)
}
func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
