package main

import "fmt"

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydog"

	byteArray := []byte(originalString)

	fmt.Println("Byte Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	selectionSort(byteArray)
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))

}
func selectionSort(s []byte) {
	var index int
	for i := 0; i < len(s); i++ {
		index = i
		current := s[i]
		for j := i; j < len(s); j++ {
			if current > s[j] {
				index = j
				current = s[j]
			}
		}
		swap(s, i, index)
	}
}

func swap(s []byte, i int, index int) {
	if i != index {
		tmp := s[i]
		s[i] = s[index]
		s[index] = tmp
	}
}
