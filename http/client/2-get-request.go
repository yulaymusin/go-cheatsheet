package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// get request is issued with the Get function
	resp, err := http.Get("http://webcode.me")
	// check for the error
	if err != nil {
		log.Fatal(err)
	}
	// client must close the response body when finished
	defer resp.Body.Close()
	// read the content of the body with ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// print the received data to the console
	fmt.Println(string(body))
}
