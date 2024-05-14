package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123454321))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12))
	fmt.Println(isPalindrome(53235))
	fmt.Println(isPalindrome(987789))
	fmt.Println(isPalindrome(9789))
}

func isPalindrome(x int) bool {
	var list = []int{}
	if x < 0 {
		return false
	}
	for x > 0 {
		if x < 10 {
			list = append(list, x)
			break
		}
		num := x % 10
		list = append(list, num)
		x = x/10
	}
	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-(i+1)] {
			return false
		}
	}
	return true
}
