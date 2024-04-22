package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloHandler).Methods("HEAD")

	log.Println("Listening...")
	http.ListenAndServe(":8080", r)
}

func HelloHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusOK)
}
