package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// map the / pattern to the HelloHandler with HandleFunc
	http.HandleFunc("/", HelloHandler)

	log.Println("Listening...")
	// ListenAndServe listens on the TCP network address and then handles requests on incoming connections
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handler responds to an HTTP request
// two parameters: the response writer and the request object
func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}
