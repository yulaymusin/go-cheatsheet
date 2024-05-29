package main

import "fmt"

func main() {
	// Basic examples
	x := 3
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("other:", x)
	}

	url := "example.com"
	switch url {
	case "example.com":
		fmt.Println("test")
	case "go.dev":
		fmt.Println("live")
	default:
		fmt.Println("dev")
	}

	// Conditional cases
	switch result := 3; {
	case result > 10:
		fmt.Println(">10")
	case result == 6:
		fmt.Println("==6")
	case result < 10:
		fmt.Println("<10")
	}

	// Case list
	switch x {
	case 1, 2, 3:
		//
	case 10, 20, 30:
		//
	}

	// Fallthrough will continue checking the next case
	letter := "a"
	switch letter {
	case " ":
	case "a", "e", "i", "o", "u":
		fmt.Println("A vowel")
		fallthrough
	case "A", "E", "I", "O", "U":
		fmt.Println("Vowels are great")
	default:
		fmt.Println("It is something else")
	}
}
