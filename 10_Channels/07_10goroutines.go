package main

import "fmt"

func main() {
	c := make(chan int)

	const (
		numGoroutines = 10
		numNumbers    = 10
	)

	for i := 0; i < numGoroutines; i++ {
		go func(sv int) {
			for j := 0; j < numNumbers; j++ {
				c <- sv + j
			}
		}(i * numNumbers)
	}

	for i := 0; i < numGoroutines*numNumbers; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("about to exit")
}
