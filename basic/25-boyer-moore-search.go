package main

import (
	"fmt"
)

func boyerMooreSearch(text string, pattern string) int {
	lastOccurrence := make([]int, 256)
	for i := 0; i < 256; i++ {
		lastOccurrence[i] = -1
	}

	patternLength := len(pattern)
	textLength := len(text)

	for i := 0; i < patternLength; i++ {
		lastOccurrence[pattern[i]] = i
	}

	i := patternLength - 1
	j := patternLength - 1

	for i < textLength {
		if text[i] == pattern[j] {
			if j == 0 {
				return i
			}
			i--
			j--
		} else {
			i += patternLength - min(j, 1+lastOccurrence[text[i]])
			j = patternLength - 1
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	text := "This is a test"
	pattern := "test"

	index := boyerMooreSearch(text, pattern)
	if index != -1 {
		fmt.Println("Pattern found at index", index)
	} else {
		fmt.Println("Pattern not found.")
	}
}
