package main

import (
	"fmt"
	"time"
)

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydog"

	byteArray := []byte(originalString)

	fmt.Println("Byte Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))

	fmt.Println("selection sort")
	startTime := time.Now()
	selectionSort(byteArray)
	endTime := time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	fmt.Println("bubbleSort")
	startTime = time.Now()
	bubbleSort(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	fmt.Println("insertionSort")
	startTime = time.Now()
	insertionSort(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	fmt.Println("shellSort")
	startTime = time.Now()
	shellSort(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

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
func selectionSortRec(s []byte, i int, j int, index int, current byte) {

}
func bubbleSort(s []byte) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				sorted = false
				swap(s, i, i+1)
			}
		}
	}
}
func insertionSort(s []byte) {
	for i := 0; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; {
			swap(s, j, j-1)
		}
	}
}

func shellSort(s []byte) {
	hSeries := [7]int{1093, 364, 121, 40, 13, 4, 1}
	for _, h := range hSeries {
		if h > len(s) {
			continue
		}
		insertionSortH(s, h)
	}

}

func insertionSortH(s []byte, h int) {
	for i := 0; i < len(s); i = i + h {
		for j := i; j > 0 && s[j] < s[j-h]; j = j - h {
			swap(s, j, j-h)
		}
	}
}
func swap(s []byte, i int, index int) {
	if i != index {
		tmp := s[i]
		s[i] = s[index]
		s[index] = tmp
	}
}
