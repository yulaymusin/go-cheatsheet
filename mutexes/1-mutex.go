package main

import (
	"fmt"
	"sync"
)

var x = 0

// increments the value of x by 1 and then calls Done() on the WaitGroup to notify its completion
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1 // this statement be executed by only one Goroutine at any point of time
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// Mutex is a struct type
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		// important to pass the address of the mutex
		go increment(&w, &m)
	}

	w.Wait()

	fmt.Println("final value of x", x)
}
