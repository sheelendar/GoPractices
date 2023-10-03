package main

import "fmt"

/*
Given Inorder Traversal of a Special Binary Tree in which the key of every node is greater than keys in left and right children, construct the Binary Tree and return root.

Examples:

Input: inorder[] = {5, 10, 28, 30, 40}
Output: root of following tree

	     28
	   /   \
	  10     30
	 /         \
	5          40
*/
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	in := []int{5, 10, 28, 30, 40}
	size := len(in)
	preIndex := 0
	node := constructBSTFromInOrder(in, 0, size-1, &preIndex)
	displayBSTInorder(node)

}

// time complexcity of this function is O(n)
func constructBSTFromInOrder(in []int, l, h int, preIndex *int) *Node {
	if l > h {
		return nil
	}
	mid := (l + h) / 2
	node := &Node{data: in[mid]}
	node.left = constructBSTFromInOrder(in, l, mid-1, preIndex)
	node.right = constructBSTFromInOrder(in, mid+1, h, preIndex)
	return node
}

func displayBSTInorder(node *Node) {
	if node == nil {
		return
	}
	displayBSTInorder(node.left)
	fmt.Print(node.data)
	fmt.Print(" ")
	displayBSTInorder(node.right)
}
