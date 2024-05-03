package main

import "fmt"

func main() {
	r := [3]int{1, 2, 3}

	if len(r) == 0 {
		fmt.Println("Empty!")
	} else {
		fmt.Println("Not empty!")
	}
}