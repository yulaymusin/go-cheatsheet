package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{
		"Zulu", "Yankee", "X-Ray", "Whiskey", "Victor",
		"Ulysses", "Tango", "Sierra", "Romeo", "Quebec",
		"Poppa", "Oscar", "Nancy", "Mike", "Lima",
		"Kilo", "Juliet", "India", "Hotel", "Golf",
		"Foxtrot", "Echo", "Delta", "Charlie", "Bravo",
		"Alpha",
	}

	fmt.Println(SortByLengthThenAlphabetically(strs))
}

func SortByLengthThenAlphabetically(strs []string) []string {
	byLengthThenAlphabetically := func(i int, j int) bool {
		x := strs[i]
		y := strs[j]
		deltaLength := len(x) - len(y)

		return deltaLength < 0 || (deltaLength == 0 && x < y)
	}
	sort.Slice(strs, byLengthThenAlphabetically)

	return strs
}
