package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// creates a GET request to a small website
	resp, err := http.Get("http://webcode.me")
	if err != nil {
		log.Fatal(err)
	}

	// Status gives the status as a string
	fmt.Println(resp.Status)
	// StatusCode gives the status as a number
	fmt.Println(resp.StatusCode)
}
