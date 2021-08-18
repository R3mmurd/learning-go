package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const max = 100
	var wg sync.WaitGroup
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // Yield the processor
			wg.Done()         // Notify to the wait group that is done
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
