package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	var counter int64

	const goRoutines = 100

	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			// Accessing the counter through the use of package atomic allows us to avoid race conditions
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))

			// Shutting down this iteration of goroutine
			wg.Done()
		}()
		fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
