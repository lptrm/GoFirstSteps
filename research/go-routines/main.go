package main

import (
	"fmt"
	"sync"
)

func sumSet(start, end int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() // decrement the wait group counter at the end of the function

	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}
	c <- sum // send the sum to the channel
}

func main() {
	const n = 100000    // number of integers to sum
	rangeSize := n / 10 // size of each range

	c := make(chan int)   // unbuffered channel
	var wg sync.WaitGroup // wait group for goroutines

	for i := 1; i < n; i += rangeSize {
		wg.Add(1)                         // add a goroutine to the wait group
		go sumSet(i, i+rangeSize, &wg, c) // start a goroutine
	}

	go func() { // anonymous goroutine for not blocking the main goroutine (in this example it's not necessary)
		wg.Wait() // wait for all goroutines to finish
		close(c)  // close the channel
	}()

	fmt.Println("Waiting for goroutines to finish...") // if the anonymous goroutine is used, this line will be printed before the goroutines finish
	result := 0
	go fmt.Println("Waiting for goroutines to finish...")
	go printX(result)
	for sum := range c { // iterates over the channel, but has to wait until the channel is closed before it can iterate
		result += sum
	}

	fmt.Println("Sum:", result)
}
func printX(y int) {
	fmt.Println(y)
}
