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
func selectionSort(arr []byte) {
	var next int
	for current := 0; current < len(arr); current++ {
		next = current
		for candidate := current + 1; candidate < len(arr); candidate++ {
			if arr[next] > arr[candidate] {
				next = candidate
			}
		}
		swap(arr, current, next)
	}
}
func selectionSortRec(s []byte, i int, j int, index int, current byte) {

}
func bubbleSort(arr []byte) {
	sorted := false
	for !sorted {
		sorted = true
		for current := 0; current < len(arr)-1; current++ {
			if arr[current] > arr[current+1] {
				sorted = false
				swap(arr, current, current+1)
			}
		}
	}
}
func insertionSort(arr []byte) {
	for current := 0; current < len(arr); current++ {
		for next := current; next > 0 && arr[next] < arr[next-1]; next-- {
			swap(arr, next, next-1)
		}
	}
}

func shellSort(arr []byte) {
	gapSeries := [7]int{1093, 364, 121, 40, 13, 4, 1}
	for _, gap := range gapSeries {
		if gap > len(arr) {
			continue
		}
		insertionSortH(arr, gap)
	}

}

func insertionSortH(arr []byte, gap int) {
	for current := 0; current < len(arr); current = current + gap {
		for next := current; next > 0 && arr[next] < arr[next-gap]; next = next - gap {
			swap(arr, next, next-gap)
		}
	}
}
func percolate(arr []byte, current, end int) {
	for next := 2 * current; next <= end; {
		if next < end && arr[next+1] > arr[next] {
			next++
		}
		if arr[next] > arr[current] {
			swap(arr, next, current)
			current = next
			next = 2 * current
		} else {
			break
		}
	}
}

func heapSort(arr []byte) {
	last := len(arr) - 1
	for current := last / 2; current >= 1; {
		percolate(arr, current, last)
		current--
	}
	for current := last; current > 1; {
		swap(arr, 1, current)
		current--
		percolate(arr, 1, current)
	}
}

func swap(arr []byte, index1 int, index2 int) {
	if index1 != index2 {
		tmp := arr[index1]
		arr[index1] = arr[index2]
		arr[index2] = tmp
	}

}
