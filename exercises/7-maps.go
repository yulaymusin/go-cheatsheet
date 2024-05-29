package main

import "fmt"

func main() {
	// Making a map
	myMap1 := make(map[string]int)
	fmt.Println(myMap1)

	myMap2 := map[string]int{
		"item 1": 1,
		"item 2": 2,
		"item 3": 3,
	}
	fmt.Println(myMap2)

	// Map operations
	// insert
	myMap1["favorite number"] = 5
	// read
	fav := myMap1["favorite number"]
	missing := myMap1["age"] // default value
	// delete
	delete(myMap1, "favorite number")
	// check existence
	price, found := myMap1["price"]
	if !found {
		fmt.Println("price not found")
	}
	fmt.Println(myMap1)
	fmt.Println(fav, missing, price)

	// Iteration
	for key, value := range myMap1 {
		_, _ = key, value
	}
}
