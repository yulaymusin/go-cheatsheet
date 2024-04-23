package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// file server is registered with Handle; it serves files from a public directory
	fileServer := http.FileServer(http.Dir("./http/server/8-file-server"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello there!\n")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
