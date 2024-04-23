package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
