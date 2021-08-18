package main

import "fmt"

func main() {
	c := make(chan int)

	// Don't! Because channels block.
	// c <- 42

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
