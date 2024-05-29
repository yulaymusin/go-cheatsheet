package main

import "fmt"

// Defining a structure
type Sample struct {
	field string
	a, b  int
}

func main() {
	// Instantiating a structure
	data1 := Sample{"word", 1, 2}
	fmt.Println(data1)

	data2 := Sample{
		field: "word",
		a:     1,
		b:     2,
	}
	fmt.Println(data2)

	// Default values
	data3 := Sample{}
	fmt.Println(data3)

	data4 := Sample{a: 5}
	fmt.Println(data4)

	// Accessing fields
	word := data1.field
	a, b := data1.a, data1.b
	fmt.Println(word, a, b)

	data1.field = "hello"
	data1.a = 10
	data1.b = 20
	fmt.Println(data1)

	// Anonymous structures
	var sample1 struct {
		field string
		a, b  int
	}
	sample1.field = "hello"
	sample1.a = 9
	fmt.Println(sample1)

	sample2 := struct {
		field string
		a, b  int
	}{
		"hello",
		1, 2,
	}
	fmt.Println(sample2)
}
