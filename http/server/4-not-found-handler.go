package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about page")
	})

	http.HandleFunc("/news", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "news page")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			w.Write([]byte("404 - not found\n"))
			return
		}

		fmt.Fprintln(w, "home page")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
