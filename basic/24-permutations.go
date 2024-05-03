package main

import "fmt"

func generatePermutations(str string) []string {
	var result []string

	if len(str) == 1 {
		return []string{str}
	}

	for i := 0; i < len(str); i++ {
		rune := str[i]

		remainingStr := str[:i] + str[i+1:]
		permutations := generatePermutations(remainingStr)

		for j := 0; j < len(permutations); j++ {
			result = append(result, string(rune)+permutations[j])
		}
	}

	return result
}

func main() {
	str := "abc"

	fmt.Println("Permutations of", str, ":", generatePermutations(str))
}
