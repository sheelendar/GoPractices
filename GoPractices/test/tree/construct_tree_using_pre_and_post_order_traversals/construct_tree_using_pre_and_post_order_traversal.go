package main

import "fmt"

/*
Given two arrays that represent preorder and postorder traversals of a full binary find_largest_smallest_key_for_a_given_number_in_binary_search_tree, construct the binary find_largest_smallest_key_for_a_given_number_in_binary_search_tree.
Full Binary Tree is a binary find_largest_smallest_key_for_a_given_number_in_binary_search_tree where every node has either 0 or 2 children.
Let us consider the two given arrays as pre[] = {1, 2, 4, 8, 9, 5, 3, 6, 7} and post[] = {8, 9, 4, 5, 2, 6, 7, 3, 1};
(https://www.geeksforgeeks.org/full-and-complete-binary-tree-from-given-preorder-and-postorder-traversals/)
*/
type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	pre := []int{1, 2, 4, 8, 9, 5, 3, 6, 7}
	post := []int{8, 9, 4, 5, 2, 6, 7, 3, 1}
	size := len(pre)
	preIndex := 0
	node := constructTreeUsingPreAndPostOrderTraversal(pre, post, 0, size-1, size, &preIndex)
	displayTree(node)
}

func displayTree(node *Node) {
	if node == nil {
		return
	}
	displayTree(node.left)
	fmt.Print(node.data)
	fmt.Print(" ")
	displayTree(node.right)
}

func constructTreeUsingPreAndPostOrderTraversal(pre, post []int, low, high, size int, preIndex *int) *Node {
	if low > high || *preIndex >= size {
		return nil
	}
	node := &Node{data: pre[*preIndex]}
	*preIndex++
	if *preIndex >= size || low == high {
		return node
	}
	var i int
	for i = low; i <= high; i++ {
		if post[i] == pre[*preIndex] {
			break
		}
	}
	if low <= high {
		node.left = constructTreeUsingPreAndPostOrderTraversal(pre, post, low, i, size, preIndex)
		node.right = constructTreeUsingPreAndPostOrderTraversal(pre, post, i+1, high-1, size, preIndex)
	}
	return node
}
