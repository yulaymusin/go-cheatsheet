package main

import "fmt"

func main() {
	// Creating a slice
	mySlice := []int{1, 2, 3}
	item1 := mySlice[0]
	fmt.Println(item1)

	numbers1 := [...]int{1, 2, 3, 4}
	slice1 := numbers1[:]   // [1, 2, 3, 4]
	slice2 := numbers1[1:]  // [2, 3, 4]
	slice3 := slice2[:1]    // [2]
	slice4 := numbers1[:2]  // [1, 2]
	slice5 := numbers1[1:3] // [2, 3]
	fmt.Println(slice1, slice2, slice3, slice4, slice5)

	// Dynamic arrays
	// The append() function can add additional elements
	numbers2 := []int{1, 2, 3}
	numbers2 = append(numbers2, 4, 5, 6)
	fmt.Println(numbers2)

	// 3 dots can be used to extend a slice with another slice
	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6}
	combined := append(part1, part2...)
	fmt.Println(combined)

	// Preallocation
	slice := make([]int, 10)
	fmt.Println(slice)

	// Slices to functions
	small := []int{1}
	big := []int{1, 2, 3, 4, 5, 6, 7}
	iterate(small)
	iterate(big)

	// Multidimensional slices
	board := [][]string{
		// type declaration is optional
		[]string{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
}

func iterate(slice []int) {
	for i := 0; i < len(slice); i++ {
		//
		fmt.Println(1)
	}
}
