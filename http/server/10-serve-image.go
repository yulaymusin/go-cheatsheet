package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// map a handler to the /image path
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/image", handler)

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}

// handler function accepts two parameters: http.ResponseWriter and http.Request
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// read the image into the buffer
	buf, err := os.ReadFile("http/server/10-server-image/image.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// set the header
	w.Header().Set("Content-Type", "image/png")
	// image data is written to the response body with Write
	w.Write(buf)
}
