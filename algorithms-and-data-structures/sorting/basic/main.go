package main

import (
	"fmt"
	"time"
)

func main() {
	originalString := " thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"

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

	byteArray = []byte(originalString)

	fmt.Println("insertionSort")
	startTime = time.Now()
	insertionSort(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	byteArray = []byte(originalString)

	fmt.Println("shellSort")
	startTime = time.Now()
	shellSort(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	byteArray = []byte(originalString)

	fmt.Println("heapSort")
	startTime = time.Now()
	heapSort(byteArray)
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
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
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

func insertionSortH(array []byte, gap int) {
	for current := 0; current < len(array); current = current + gap {
		for next := current; next > 0 && array[next] < array[next-gap]; next = next - gap {
			swap(array, next, next-gap)
		}
	}
}
func percolate(a []byte, j, t int) {
	h := 2 * j // when putting this in the for loop, it wont work???
	for h <= t {
		if h < t && a[h+1] > a[h] {
			h++
		}
		if a[h] > a[j] {
			swap(a, h, j)
			j = h
			h = 2 * j
		} else {
			break
		}
	}
}

func heapSort(a []byte) {
	hi := len(a) - 1
	j := len(a) - 1
	for j = hi / 2; j >= 1; {
		percolate(a, j, hi)
		j--
	}
	for j = hi; j > 1; {
		swap(a, 1, j)
		j--
		percolate(a, 1, j)
	}
}

func swap(array []byte, index1 int, index2 int) {
	if index1 != index2 {
		tmp := array[index1]
		array[index1] = array[index2]
		array[index2] = tmp
	}

}
