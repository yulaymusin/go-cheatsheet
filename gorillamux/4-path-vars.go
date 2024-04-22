package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello/{name}", func(resp http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		name := vars["name"]

		fmt.Fprintf(resp, "Hello %s!", name)
	})

	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}
