package main

import (
	"fmt"
	"time"
)

func main() {
	originalString := "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"

	byteArray := []byte(originalString)

	fmt.Println("radixSort")
	startTime := time.Now()
	radixSort(byteArray)
	endTime := time.Now()
	fmt.Println("Sorted Array:", byteArray)
	fmt.Println("String Rep:", string(byteArray))
	fmt.Println("The Time was:", endTime.Sub(startTime))

}
func radixSort(arr []byte) {
}
