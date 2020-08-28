package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	counter := 0
   
   	const goRoutines = 100

	wg.Add(goRoutines)
	
	for i := 0; i < goRoutines; i++ {
		go func() {
			v := counter
			runtime.Gosched()	
			v++
			counter = v
			wg.Done()
		}()
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
   	}
	
	wg.Wait()
	
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("count:\t", counter)
}