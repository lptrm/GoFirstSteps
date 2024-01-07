package main

import (
	"fmt"
	"time"
)

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"
	byteArray := []byte(originalString)

	fmt.Println("mergeSort")
	startTime := time.Now()
	mergeSort(byteArray)
	endTime := time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))
}
func mergeSort(arr []byte) {
	mergeSortR(arr, 0, len(arr)-1)
}
func mergeSortR(arr []byte, start int, end int) {
	if start < end {
		mid := (start + end + 1) / 2
		mergeSortR(arr, start, mid-1)
		mergeSortR(arr, mid, end)
		tmp := make([]byte, end-start+1)
		left := start
		right := mid
		for current := 0; current < len(tmp); current++ {
			if (right > end) || (left < mid) && (arr[left] < arr[right]) {
				tmp[current] = arr[left]
				left++
			} else {
				tmp[current] = arr[right]
				right++
			}
		}
		for i := 0; i < len(tmp); i++ {
			arr[start+i] = tmp[i]
		}
	}
}
