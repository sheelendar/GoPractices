package main

import "fmt"

/*
Given postorder traversal of a binary search tree, construct the BST.

For example, if the given traversal is {1, 7, 5, 50, 40, 10}, then following tree should be constructed and root of the tree should be returned.

		    10
		  /   \
		 5     40
		/  \      \
	   1    7      50
*/
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	pre := []int{1, 7, 5, 50, 40, 10}
	size := len(pre)
	postIndex := size - 1
	node := constructBSTuisngPostOrder(pre, -999999, 999999, &postIndex)
	displayBSTInorder(node)

}

// time complexcity of this function is O(n)
func constructBSTuisngPostOrder(pre []int, min int, max int, postIndex *int) *Node {
	if *postIndex < 0 || min > pre[*postIndex] || pre[*postIndex] > max {
		return nil
	}
	node := &Node{data: pre[*postIndex]}
	*postIndex--
	node.right = constructBSTuisngPostOrder(pre, node.data, max, postIndex)
	node.left = constructBSTuisngPostOrder(pre, min, node.data, postIndex)
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
