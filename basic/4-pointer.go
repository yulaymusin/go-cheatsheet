package main

import "fmt"

func main() {
	var x = 5
	var p *int
	p = &x

	fmt.Printf("x = %d\n", *p)
}
