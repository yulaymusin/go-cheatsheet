package main

import (
	"fmt"
	"sort"
)

func main() {
	// init slices of strings
	scl1 := []string{"abc", "rwp", "def", "por", "ber", "erj"}
	scl2 := []string{"Rabbit", "Fish", "Dog", "Parrot", "Cat", "Hamster"}

	// display content of the slices
	fmt.Println("Slices (Before sort):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	// check whether the slices is in sorted form (in increasing order) or not
	res1 := sort.StringsAreSorted(scl1)
	res2 := sort.StringsAreSorted(scl2)
    fmt.Println("Is Slice 1 is sorted?: ", res1) 
    fmt.Println("Is Slice 2 is sorted?: ", res2) 
	
	// sort the slices
	sort.Strings(scl1)
	sort.Strings(scl2)

	// display content of the slices
	fmt.Println("Slices (After sort):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	// check again whether the slices is in sorted form (in increasing order) or not
	res1 = sort.StringsAreSorted(scl1)
	res2 = sort.StringsAreSorted(scl2)
    fmt.Println("Is Slice 1 is sorted?: ", res1) 
    fmt.Println("Is Slice 2 is sorted?: ", res2) 
}
