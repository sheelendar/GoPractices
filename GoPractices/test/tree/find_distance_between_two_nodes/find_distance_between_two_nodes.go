package main

import (
	"fmt"
)

/*
AucFind the distance between two keys in a binary find_largest_smallest_key_for_a_given_number_in_binary_search_tree, no parent pointers are given.
The distance between two nodes is the minimum number of edges to be traversed to reach one node from another.
(https://www.geeksforgeeks.org/find-distance-between-two-nodes-of-a-binary-tree/)
*/
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	node := &Node{data: 1}
	node.left = &Node{data: 2}
	node.right = &Node{data: 3}

	node.left.left = &Node{data: 4}
	node.left.right = &Node{data: 5}
	node.right.left = &Node{data: 6}
	node.right.right = &Node{data: 7}
	node.right.left.right = &Node{data: 8}
	fmt.Println("Dist(4, 5) = ", findDistanceBWNodes(node, 4, 5))
	fmt.Println("Dist(4, 6) = ", findDistanceBWNodes(node, 6, 4))
	fmt.Println("Dist(3, 4) = ", findDistanceBWNodes(node, 3, 4))
	fmt.Println("Dist(2, 4) = ", findDistanceBWNodes(node, 2, 4))
	fmt.Println("Dist(8, 5) = ", findDistanceBWNodes(node, 8, 5))
}

func findDistanceBWNodes(node *Node, n1 int, n2 int) int {
	if node == nil {
		return -1
	}
	cn := findCommonNode(node, n1, n2)
	d1 := findDistanceFromCommonNode(cn, n1, 0)
	d2 := findDistanceFromCommonNode(cn, n2, 0)
	return d1 + d2
}

func findDistanceFromCommonNode(cn *Node, n, level int) int {
	if cn == nil {
		return -1
	}
	if cn.data == n {
		return level
	}
	l := findDistanceFromCommonNode(cn.left, n, level+1)
	if l != -1 {
		return l
	}
	return findDistanceFromCommonNode(cn.right, n, level+1)
}

func findCommonNode(node *Node, n1 int, n2 int) *Node {
	if node == nil {
		return nil
	}
	if node.data == n1 || node.data == n2 {
		return node
	}
	left := findCommonNode(node.left, n1, n2)
	right := findCommonNode(node.right, n1, n2)

	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
