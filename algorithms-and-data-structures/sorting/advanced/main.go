package main

import (
	"awesomeProject/algorithms-and-data-structures/sorting"
	"fmt"
	"time"
)

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"
	byteArray := []byte(originalString)

	fmt.Println("heapSort")
	startTime := time.Now()
	heapSort(byteArray)
	endTime := time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

	byteArray = []byte(originalString)

	fmt.Println("heapSortButtomUp")
	startTime = time.Now()
	heapSortButtomUp(byteArray)
	endTime = time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))
}
func heapify(arr []byte) {
	last := len(arr) - 1
	for current := last / 2; current >= 1; {
		percolate(arr, current, last)
		current--
	}
}

func percolate(arr []byte, current, end int) {
	for next := 2 * current; next <= end; {
		if next < end && arr[next+1] > arr[next] {
			next++
		}
		if arr[next] > arr[current] {
			sorting.Swap(arr, next, current)
			current = next
			next = 2 * current
		} else {
			break
		}
	}
}

func heapSort(arr []byte) {
	last := len(arr) - 1
	heapify(arr)
	for current := last; current > 1; {
		sorting.Swap(arr, 1, current)
		current--
		percolate(arr, 1, current)
	}

}
func heapSortButtomUp(arr []byte) {
	last := len(arr) - 1
	heapify(arr)
	for current := last; current > 1; {
		sorting.Swap(arr, 1, current)
		current--
		percolateButtomUp(arr, 1, current)
	}
}
func percolateButtomUp(arr []byte, current, end int) {
	for next := 2 * current; next <= end; {
		if next < end && arr[next+1] > arr[next] {
			next++
		}
		sorting.Swap(arr, next, current)
		current = next
		next = 2 * current
	}
	percolateUp(arr, current)
}
func percolateUp(arr []byte, current int) {
	for next := current / 2; next > 0; {
		if arr[next] < arr[current] {
			sorting.Swap(arr, next, current)
			current = next
			next = current / 2
		} else {
			break
		}
	}
}
