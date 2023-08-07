package main

import "fmt"

/*
Consider lines of slope -1 passing between nodes (dotted lines in below diagram).
The diagonal sum in a binary find_largest_smallest_key_for_a_given_number_in_binary_search_tree is the sum of all node’s data lying between these lines.
Given a Binary Tree, print all diagonal sums.
(https://www.geeksforgeeks.org/diagonal-sum-binary-tree/)
For the following input find_largest_smallest_key_for_a_given_number_in_binary_search_tree, the output should be 9, 19, 42.
9 is sum of 1, 3 and 5.
19 is sum of 2, 6, 4 and 7.
42 is sum of 9, 10, 11 and 12.
*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 9}
	root.left.right = &Node{data: 6}
	root.right.left = &Node{data: 4}
	root.right.right = &Node{data: 5}
	root.right.left.left = &Node{data: 12}
	root.right.left.right = &Node{data: 7}
	root.left.right.left = &Node{data: 11}
	root.left.left.right = &Node{data: 10}
	m := make(map[int]int, 100)
	getDiagonalSum(root, 0, m)
	for _, v := range m {
		fmt.Print(v)
		fmt.Print(" ")
	}
}

func getDiagonalSum(root *Node, level int, m map[int]int) {
	if root == nil {
		return
	}
	if v, ok := m[level]; ok {
		m[level] = v + root.data
	} else {
		m[level] = root.data
	}
	getDiagonalSum(root.left, level+1, m)
	getDiagonalSum(root.right, level, m)
}
