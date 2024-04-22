package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// /now path is mapped to the NowHandler function
	r.HandleFunc("/now", NowHandler)

	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}

func NowHandler(resp http.ResponseWriter, _ *http.Request) {
	// get the current datetime with time.Now
	now := time.Now()

	// add the formatted datetime to a Go map
	payload := make(map[string]string)
	payload["now"] = now.Format(time.ANSIC)

	// set the appropriate header of the response
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	// encoded payloaded is sent back to the client
	json.NewEncoder(resp).Encode(payload)
}
