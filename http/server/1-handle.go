package main

import (
	"fmt"
	"net/http"
)

// CounterHandler contains the counter variable and the implementation of the ServeHTTP function
type CounterHandler struct {
	counter int
}

// increments the counter and writes a message to the http.ResponseWriter
func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++
	fmt.Fprintln(w, "Counter:", ct.counter)
}

func main() {
	// CounterHandler is created and registered with Handle
	th := &CounterHandler{counter: 0}
	http.Handle("/count", th)

	http.ListenAndServe(":8080", nil)
}
