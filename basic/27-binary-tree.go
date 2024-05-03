package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{data: data, left: nil, right: nil}
	}

	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

func search(root *Node, data int) bool {
	if root == nil {
		return false
	}

	if root.data == data {
		return true
	}

	if data < root.data {
		return search(root.left, data)
	}

	return search(root.right, data)
}

func main() {
	root := &Node{}

	root = insert(root, 50)
	root = insert(root, 30)
	root = insert(root, 20)
	root = insert(root, 40)
	root = insert(root, 70)
	root = insert(root, 60)
	root = insert(root, 80)

	fmt.Println("Searching for 40:", search(root, 40))
	fmt.Println("Searching for 100:", search(root, 100))
}