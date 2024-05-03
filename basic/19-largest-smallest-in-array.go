package main

import "fmt"

func main() {
	arr := []int{5, 8, 2, 10, 3}

	max := arr[0]
	min := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Println("Largest element:", max)
	fmt.Println("Smallest element:", min)
}
