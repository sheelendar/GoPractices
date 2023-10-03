package main

import "fmt"

/*
Given the preorder traversal of a binary search find_largest_smallest_key_for_a_given_number_in_binary_search_tree, construct the BST.

Examples:

		 Input: {10, 5, 1, 7, 40, 50}
		Output:          10
		               /   \
		              5     40
	                /  \       \
	              1    7       50
*/
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	pre := []int{10, 5, 1, 7, 40, 50}
	size := len(pre)
	preIndex := 0
	//node := constructBSTFromPreorder(pre, 0, size-1, size, &preIndex)
	node := constructBSTUsingMinAndMax(pre, -999999, 999999, size, &preIndex)
	displayBSTInorder(node)

}

// time complexcity of this function is O(n)
func constructBSTUsingMinAndMax(pre []int, min int, max int, size int, preIndex *int) *Node {
	if *preIndex >= size || min > pre[*preIndex] || pre[*preIndex] > max {
		return nil
	}
	node := &Node{data: pre[*preIndex]}
	*preIndex++

	node.left = constructBSTUsingMinAndMax(pre, min, node.data, size, preIndex)
	node.right = constructBSTUsingMinAndMax(pre, node.data, max, size, preIndex)
	return node
}

// time complexcity of this function is O(n2)
func constructBSTFromPreorder(pre []int, low int, high int, size int, preIndex *int) *Node {
	if *preIndex >= size || low > high {
		return nil
	}
	node := &Node{data: pre[*preIndex]}
	*preIndex++
	if low == high || *preIndex >= size {
		return node
	}
	var i int
	for i = low; i < high; i++ {
		if pre[i] > node.data {
			break
		}
	}
	if low < high {
		node.left = constructBSTFromPreorder(pre, *preIndex, i-1, size, preIndex)
		node.right = constructBSTFromPreorder(pre, i, high, size, preIndex)
	}
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
