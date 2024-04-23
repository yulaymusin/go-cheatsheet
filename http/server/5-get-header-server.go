package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ua", func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		fmt.Fprintf(w, "User agent: %s\n", ua)
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
