package main

import (
	"fmt"
	"sync"
)

var GFG = 0

// will runs in every goroutine, a WaitGroup must be passed to functions by pointer
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	// Lock() the mutex to ensure exclusive access to the state
	m.Lock()
	GFG = GFG + 1
	m.Unlock()

	// on return, notify the WaitGroup that we’re done
	wg.Done()
}
func main() {
	// WaitGroup is used to wait for all the goroutines launched here to finish
	var w sync.WaitGroup

	// This mutex will synchronize access to state.
	var m sync.Mutex

	// launch several goroutines and increment the WaitGroup counter for each
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w, &m)
	}
	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done
	w.Wait()
	fmt.Println("Value of x", GFG)
}
