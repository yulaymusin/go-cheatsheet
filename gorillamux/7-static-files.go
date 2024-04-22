package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// http.FileServer serves static data from a public directory
	// URL paths starting with /app/ are directed to this directory
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir("./gorillamux/7-static-files"))))

	log.Println("Listening...")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
