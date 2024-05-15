package main

import "fmt"

func main() {
	strs1 := []string{"flower", "flow", "flight"} // 4
	strs2 := []string{"Zulu", "Yankee", "X-Ray", "Whiskey", "Victor"} // 4
	strs3 := []string{"Ulysses", "Tango", "Sierra", "Romeo", "Quebec"} // 5
	strs4 := []string{"Poppa", "Oscar", "Nancy", "Mike", "Lima"} // 4
	strs5 := []string{"Kilo", "Juliet", "India", "Hotel", "Golf"} // 4
	strs6 := []string{"Foxtrot", "Echo", "Delta", "Charlie", "Bravo"} // 4
	strs7 := []string{"Alpha"} // 5

	fmt.Println(ShortestLen(strs1))
	fmt.Println(ShortestLen(strs2))
	fmt.Println(ShortestLen(strs3))
	fmt.Println(ShortestLen(strs4))
	fmt.Println(ShortestLen(strs5))
	fmt.Println(ShortestLen(strs6))
	fmt.Println(ShortestLen(strs7))
}
func ShortestLen(strs []string) int {
	charCount := 0
	for i := 0; i < len(strs); i++ {
		if i == 0 {
			charCount = len(strs[0])
		}
		if i > 0 && len(strs[i]) < charCount {
			charCount = len(strs[i])
		}
	}
	return charCount
}
