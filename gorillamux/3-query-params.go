package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		// get the name value from the request
		name := req.URL.Query().Get("name")

		if name == "" {
			name = "guest"
		}

		fmt.Fprintf(resp, "Hello %s!", name)
	})

	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}
