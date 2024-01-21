package main

import (
	"awesomeProject/algorithms-and-data-structures/sorting"
	"fmt"
	"time"
)

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"

	byteArray := []byte(originalString)
	fmt.Println("quickSort")
	startTime := time.Now()
	quickSort(byteArray)
	endTime := time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

}
func quickSort(arr []byte) {
	last := len(arr) - 1
	quickSortR(arr, 0, last)
}
func quickSortR(arr []byte, start int, end int) {
	left := start
	right := end
	mid := (left + right) / 2
	if arr[left] > arr[mid] {
		sorting.Swap(arr, left, mid)
	}
	if arr[mid] > arr[right] {
		sorting.Swap(arr, mid, right)
	}
	if arr[left] > arr[mid] {
		sorting.Swap(arr, left, mid)
	}
	if right-left > 2 {
		pivot := arr[mid]
		for left <= right {
			for arr[left] < pivot {
				left++
			}
			for arr[right] > pivot {
				right--
			}
			if left <= right {
				sorting.Swap(arr, left, right)
				left++
				right--
			}
		}
		if start < right {
			quickSortR(arr, start, right)
		}
		if left < end {
			quickSortR(arr, left, end)
		}
	}
}
