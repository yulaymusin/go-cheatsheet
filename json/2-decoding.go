package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {
	var u1 User

	// JSON object which is transformed into bytes
	data := []byte(`{
        "Id" : 1,
        "Name": "John Doe",
        "Occupation": "gardener"
    }`)

	// decode the u1 struct from JSON into a struct with Unmarshal
	err := json.Unmarshal(data, &u1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Struct is:", u1)
	fmt.Printf("%s is a %s.\n", u1.Name, u1.Occupation)

	// declare a slice of User structs
	var u2 []User

	// JSON array to be decoded into a Go slice
	data2 := []byte(`
    [
        {"Id":2,"Name":"Roger Roe","Occupation":"driver"},
        {"Id":3,"Name":"Lucy Smith","Occupation":"teacher"},
        {"Id":4,"Name":"David Brown","Occupation":"programmer"}
    ]`)

	// decode the JSON array to Go slice with Unmarshal
	err2 := json.Unmarshal(data2, &u2)

	if err2 != nil {
		log.Fatal(err2)
	}

	// print the decoded JSON array user by user
	for i := range u2 {
		fmt.Println(u2[i])
	}
}
