package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return item, nil
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	item, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(item)
	}
}