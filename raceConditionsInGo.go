package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	counter := 0

	const goRoutines = 100

	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			// The mutex lock allows the 'counter' variable to be locked while this block of code is executing.
			// This effectively eliminates the undesirable results we get from race conditions.
			mu.Lock()

			v := counter

			// The following line yields other concurrent funcs to be executed
			runtime.Gosched()

			v++
			counter = v

			// Unlocking the counter variable
			mu.Unlock()

			// Shutting down this iteration of goroutine
			wg.Done()
		}()
		fmt.Println("Goroutines:\t", runtime.NumGoroutine())
		fmt.Println("count:\t", counter)
	}

	wg.Wait()

	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}
