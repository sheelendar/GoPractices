package main

import "fmt"

func main() {
	subRoot := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	root := &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}
	root.Left = subRoot
	fmt.Print(isSubtree(root, subRoot))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}
	if root.Val == subRoot.Val && helper(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func helper(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}
	return root.Val == subRoot.Val && helper(root.Left, subRoot.Left) && helper(root.Right, subRoot.Right)
}
