package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	N := buildTree(preorder, inorder)
	preOrder(N)
}

func preOrder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Val)
	fmt.Print(" ")
	preOrder(n.Left)
	preOrder(n.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil || inorder == nil {
		return nil
	}
	index := 0
	l := 0
	h := len(preorder)
	return bt(preorder, inorder, &index, l, h)
}

func bt(pre []int, in []int, index *int, l, h int) *TreeNode {
	if l > h {
		return nil
	}
	N := &TreeNode{Val: pre[*index]}
	*index++
	if l == h {
		return N
	}
	if *index == len(pre) {
		return N
	}
	i := l
	for ; i < h; i++ {
		if N.Val == in[i] {
			break
		}
	}

	N.Left = bt(pre, in, index, l, i-1)
	N.Right = bt(pre, in, index, i+1, h)
	return N
}
