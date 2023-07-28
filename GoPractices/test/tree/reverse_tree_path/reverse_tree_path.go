package main

import "fmt"

/*

Given a tree and node data, the task to reverse the path to that particular Node.

Examples:

Input:
            7
         /    \
        6       5
       / \     / \
      4  3     2  1
Data = 4
Output: Inorder of tree
7 6 3 4 2 5 1


Input:
            7
         /    \
        6       5
       / \     / \
      4  3     2  1
Data = 2
Output : Inorder of tree
4 6 3 2 7 5 1
*/

type Node struct {
	Data  int
	left  *Node
	right *Node
}

func main() {
	root := Node{Data: 7}
	root.left = &Node{Data: 6}
	root.right = &Node{Data: 5}
	root.left.left = &Node{Data: 4}
	root.left.right = &Node{Data: 3}
	root.right.left = &Node{Data: 2}
	root.right.right = &Node{Data: 1}
	dp := make(map[int]int)
	index := 0
	position := 0
	data := 4
	inorder(&root)
	reverseTreePath(&root, data, dp, index, position)
	fmt.Println()
	inorder(&root)
}

func inorder(root *Node) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Print(root.Data)
	fmt.Print(" ")
	inorder(root.right)
}

func reverseTreePath(root *Node, data int, dp map[int]int, level int, position int) *Node {
	if root == nil {
		return nil
	}
	if root.Data == data {
		root.Data = dp[position]
		dp[position] = data
		position++
		return root
	}
	dp[level] = root.Data
	var right *Node
	left := reverseTreePath(root.left, data, dp, level+1, position)
	if left == nil {
		right = reverseTreePath(root.left, data, dp, level+1, position)
	}
	if left != nil || right != nil {
		root.Data = dp[position]
		position++
	}
	if left != nil {
		return left
	}
	return right
}
