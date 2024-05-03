package main

import "fmt"

func main() {
	a := map[string]bool{"A": true, "B": true}
	b := make(map[string]bool)
	for k, v := range a {
		b[k] = v
	}
	fmt.Println(a, b)
}
