package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Count the number of words in this sentence."
	words := strings.Split(text, " ")
	fmt.Println("Number of words:", len(words))
}
