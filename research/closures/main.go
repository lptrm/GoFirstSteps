package main

import (
	"fmt"
	"time"
)

func main() {
	weapon := "Katana"

	go fmt.Println("Armed with", weapon)

	go fmt.Println("Armed with", weapon)

	time.Sleep(1 * time.Second)
}
