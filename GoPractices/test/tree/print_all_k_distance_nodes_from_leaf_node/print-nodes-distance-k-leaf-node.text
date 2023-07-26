package main

import (
	"fmt"
)

/*
Given a Binary Tree and a positive integer K, print all nodes that are distance K from a leaf node.
Here K distance from a leaf means K levels higher than a leaf node. For example,
if K is more than the height of the Binary Tree, then nothing should be printed.
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
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}
	root.right.left.right = &Node{data: 8}
	k := 2
	level := 0
	path := make([]int, 1000)
	visited := make([]bool, 1000)
	//printKDistanceNode(root, k, level)
	printKDistanceNodes(root, k, level, path, visited)
}
func printKDistanceNodes(root *Node, k, level int, path []int, visited []bool) {
	if root == nil {
		return
	}
	path[level] = root.data
	visited[level] = false
	level++
	if root.left == nil && root.right == nil && level-k-1 >= 0 && !visited[level-k-1] {
		fmt.Print(path[level-k-1])
		fmt.Print(" ")
		visited[level-k-1] = true
	}
	printKDistanceNodes(root.left, k, level, path, visited)
	printKDistanceNodes(root.right, k, level, path, visited)
}

func printKDistanceNode(root *Node, k, level int) []int {
	if root == nil {
		return []int{0}
	}
	if root.left == nil && root.right == nil {
		return []int{1}
	}
	leftC := printKDistanceNode(root.left, k, level)
	for i := 0; i < len(leftC); i++ {
		if leftC[i] == k {
			fmt.Print(root.data)
			fmt.Print(" ")
		}
	}

	rightC := printKDistanceNode(root.right, k, level)
	for i := 0; i < len(rightC); i++ {
		if rightC[i] == k {
			fmt.Print(root.data)
			fmt.Print(" ")
		}
	}
	var list []int
	for i := 0; i < len(leftC); i++ {
		list = append(leftC, leftC[i]+1)
	}
	for i := 0; i < len(rightC); i++ {
		list = append(leftC, rightC[i]+1)
	}
	return list
}
