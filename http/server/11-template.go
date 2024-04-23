package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name       string
	Occupation string
}

type Data struct {
	Users []User
}

func main() {
	tmp := template.Must(template.ParseFiles("http/server/11-template/layout.html"))

	http.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		data := Data{
			Users: []User{
				{Name: "John Doe", Occupation: "gardener"},
				{Name: "Roger Roe", Occupation: "driver"},
				{Name: "Peter Smith", Occupation: "teacher"},
			},
		}
		tmp.Execute(w, data)
	})

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
