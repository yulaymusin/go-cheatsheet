package main

import (
	"fmt"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		// send the form.html to the client
		http.ServeFile(w, r, "http/server/9-get-post/form.html")
	case "POST":
		// parses the raw query from the URL and updates r.Form
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// get the form values with FormValue
		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", process)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
