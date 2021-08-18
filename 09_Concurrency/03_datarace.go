package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Prove that it is a race condition by using the -race flag

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0

	const max = 100

	var wg sync.WaitGroup
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			x := counter      // Read the counter
			runtime.Gosched() // Yield the processor
			x++               // Increment the new variable
			counter = x       // Write the new value to the counter
			wg.Done()         // Notify to the wait group that is done
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
