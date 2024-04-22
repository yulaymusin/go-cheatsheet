package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// HandleFunc registers a new route with a matcher for the URL path
	r.HandleFunc("/hello", func(resp http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(resp, "Hello there!")
	})

	log.Println("Listening...")

	// pass the router to the ListenAndServe function
	http.ListenAndServe(":8080", r)
}
