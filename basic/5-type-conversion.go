package main

import "fmt"

func main() {
	var i int = 55
	var j float64 = 67.8

	fmt.Println(i + int(j))
	fmt.Println(float64(i) + j)
}
