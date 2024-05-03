package main

import "fmt"

func longestIncreasingSubsequence(arr []int) []int {
	n := len(arr)

	lengths := make([]int, n)
	sequences := make([][]int, n)

	for i := 0; i < n; i++ {
		lengths[i] = 1
		sequences[i] = []int{arr[i]}

		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lengths[j]+1 > lengths[i] {
				lengths[i] = lengths[j] + 1
				sequences[i] = append(sequences[j], arr[i])
			}
		}
	}

	maxLength := 0
	maxSequence := []int{}

	for i := 0; i < n; i++ {
		if lengths[i] > maxLength {
			maxLength = lengths[i]
			maxSequence = sequences[i]
		}
	}

	return maxSequence
}

func main() {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60}

	fmt.Println("Longest increasing subsequence:", longestIncreasingSubsequence(arr))
}