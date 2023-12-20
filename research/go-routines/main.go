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
	const n = 100       // number of integers to sum
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

	result := 0
	for sum := range c { // iterate over the channel
		result += sum
	}

	fmt.Println("Sum:", result)
}