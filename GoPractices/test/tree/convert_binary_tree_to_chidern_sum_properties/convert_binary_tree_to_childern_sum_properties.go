package main

import "fmt"

/*
Convert an arbitrary Binary Tree to a find_largest_smallest_key_for_a_given_number_in_binary_search_tree that holds Children Sum Property
*/
/*          50
         /     \
       /         \
     7             2
   / \             /\
 /     \          /   \
3        5      1      30

*/

type TreeNode struct {
	data  int64
	left  *TreeNode
	right *TreeNode
}

func main() {
	node := &TreeNode{data: 50}
	node.left = &TreeNode{data: 7}
	node.right = &TreeNode{data: 2}
	node.left.left = &TreeNode{data: 3}
	node.left.right = &TreeNode{data: 5}
	node.right.left = &TreeNode{data: 1}
	node.right.right = &TreeNode{data: 30}
	inorderTraversal(node)
	fmt.Println()
	convertChildrenSumProperties(node)
	inorderTraversal(node)
}

func convertChildrenSumProperties(node *TreeNode) {
	if node == nil || (node.left == nil && node.right == nil) {
		return
	} else {
		convertChildrenSumProperties(node.left)
		convertChildrenSumProperties(node.right)
		leftSum := int64(0)
		rightSum := int64(0)
		if node.left != nil {
			leftSum = node.left.data
		}
		if node.right != nil {
			rightSum = node.right.data
		}
		sumDiff := leftSum + rightSum - node.data

		if sumDiff > 0 {
			node.data = node.data + sumDiff
		}
		if sumDiff < 0 {
			incrementChields(node, -sumDiff)
		}
	}
}

func incrementChields(node *TreeNode, diff int64) {
	if node.left != nil {
		node.left.data = node.left.data + diff
		incrementChields(node.left, diff)
	} else if node.right != nil {
		node.right.data = node.right.data + diff
		incrementChields(node.right, diff)
	}

}

func inorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraversal(node.left)
	fmt.Print(node.data)
	fmt.Print(" ")
	inorderTraversal(node.right)
}

func convertChildrenSumProperties2(node *TreeNode) {
	if node == nil {
		return
	}
	chieldSum := int64(0)
	if node.left != nil {
		chieldSum = chieldSum + node.left.data
	}
	if node.right != nil {
		chieldSum = chieldSum + node.right.data
	}
	if chieldSum > node.data {
		node.data = chieldSum
	} else {
		if node.left != nil {
			node.left.data = node.data

		}
		if node.right != nil {
			node.right.data = node.data

		}
	}
	convertChildrenSumProperties2(node.left)
	convertChildrenSumProperties2(node.right)
	chldSum := int64(0)

	if node.left != nil {
		chldSum = chldSum + node.left.data
	}
	if node.right != nil {
		chldSum = chldSum + node.right.data
	}
	if node.left != nil || node.right != nil {
		node.data = chldSum
	}

}
