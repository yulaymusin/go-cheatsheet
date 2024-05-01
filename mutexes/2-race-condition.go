package main

import (
	"fmt"
	"sync"
)

var GFG = 0

// will runs in every goroutine, a WaitGroup must be passed to functions by pointer
func worker(wg *sync.WaitGroup) {
	GFG = GFG + 1

	// on return, notify the WaitGroup that we’re done
	wg.Done()
}
func main() {
	// WaitGroup is used to wait for all the goroutines launched here to finish
	var w sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}
	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done
	w.Wait()
	fmt.Println("Value of x", GFG)
}
