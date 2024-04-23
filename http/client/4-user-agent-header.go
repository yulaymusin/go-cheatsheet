package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// create the http client with a 3 s timeout
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	// new request is created with http.NewRequest
	req, err := http.NewRequest("GET", "http://webcode.me/ua.php", nil)
	if err != nil {
		log.Fatal(err)
	}

	// add the User-Agent header to the request
	req.Header.Add("User-Agent", "Go program")
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// requested resource simply returns the client's User-Agent string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
