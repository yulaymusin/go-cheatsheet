package main

import (
	"fmt"
	"sync"
)

var x = 0

// increments the value of x by 1 and then calls Done() on the WaitGroup to notify its completion
func increment(wg *sync.WaitGroup) {
	// race condition occurs when trying to increment x
	x = x + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// spawn 10000 increment Goroutines run concurrently
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()

	fmt.Println("final value of x", x)
}
