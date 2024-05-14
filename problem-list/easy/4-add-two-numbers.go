package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l1 := []int{}
	l2 := []int{}

	l1 = []int{2, 4, 3}
	l2 = []int{5, 6, 4}
	fmt.Println(addTwoNumbers(l1, l2)) // [7 0 8]

	l1 = []int{0}
	l2 = []int{0}
	fmt.Println(addTwoNumbers(l1, l2)) // [0]

	l1 = []int{9,9,9,9,9,9,9}
	l2 = []int{9,9,9,9}
	fmt.Println(addTwoNumbers(l1, l2)) // [8 9 9 9 0 0 0 1]
}

func addTwoNumbers(l1 []int, l2 []int) []int {
	funcToInt := func(l []int) int {
		// make a slice of strings from the slice of ints
		ls := []string{}
		for i := 0; i < len(l); i++ {
			ls = append(ls, strconv.Itoa(l[i]))
		}
		// make a mirrored slice of strings
		mls := []string{}
		for i := len(ls) - 1; i >= 0; i-- {
			mls = append(mls, ls[i])
		}
		// join a slice os strings into a single string
		s := strings.Join(mls[:], "")
		// convert string to int
		n, _ := strconv.Atoi(s)

		return n
	}
	_ = funcToInt

	n1 := funcToInt(l1)
	n2 := funcToInt(l2)
	theSumInt := n1 + n2
	theSumString := strconv.Itoa(theSumInt)   // convert an int to string
	theSum := strings.Split(theSumString, "") // split string
	// make a slice of ints
	resultSum := []int{}
	for i := len(theSum) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(theSum[i])
		resultSum = append(resultSum, v)
	}

	return resultSum
}
