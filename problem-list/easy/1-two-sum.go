package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var a1 int
	var a2 int
	for i := 0; i < len(nums); i++ {
		for ni := 0; ni < len(nums); ni++ {
			if i != ni {
				if (nums[i] + nums[ni]) == target {
					a1 = i
					a2 = ni
					break
				}
			}
		}
	}
	return []int{a1, a2}
}

func main() {
	nums := []int{2,7,11,15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
