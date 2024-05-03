package main

import "fmt"

func main() {
	// The for-loop syntax:
	// for initialization; condition; postcondition {
	// 	do something
	// }
	for i := 0; i < 3; i++ { // start of the execution block
		fmt.Println("Hello") // prints "Hello" 3 times
	} // end

	// The infinite for-loop:
	// for {
	// 	this block executes infinitely
	// }

	// The conditional for-loop:
	i := 0
	for i < 5 { // it checks the condition and runs if it matches
		fmt.Println(i) // prints 0 2 4
		i += 2
	}

	// The range-for loop:
	var items []int = []int{1, 2, 3, 4, 5}
	for i, v := range items { // iterate over a slice or array of items
		fmt.Println(i, v)
	}
	// if the current value is the only thing needed then ignore the index using the blank identifier
	for _, v := range items {
		fmt.Println(v)
	}

	// The range-for with maps:
	var m map[int]string = map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}
	for k, v := range m { // range-for can be used with maps which will provide the key and value both at the same time
		fmt.Println(k, v)
	}

	// The nested-for loop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}

	// Using labels with the for-loop:
outside: // declare the label
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i < j {
				break outside // break to that label
			}
			fmt.Println(i, j)
		}
	}
	// use labels to break out of an ongoing loop

	// The continue statement in for-loop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i < j {
				continue                           // forces to move to the next iteration
				fmt.Println("will not be printed") // skip the segment of the code
			}
			fmt.Println(i, j)
		}
	}

	// Increment
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
}
