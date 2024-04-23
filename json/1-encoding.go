package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// declare the User struct
type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {
	// create the struct instance
	u1 := User{1, "John Doe", "gardener"}

	// encode the u1 struct into JSON with Marshal
	json_data, err := json.Marshal(u1)

	if err != nil {
		log.Fatal(err)
	}

	// json_data is a byte array, we convert it to a string with the string function
	fmt.Println(string(json_data))

	// slice of users
	users := []User{
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
	}

	// encode the slice of users with Marshal
	json_data2, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	// print the encoded slice
	fmt.Println(string(json_data2))
}
