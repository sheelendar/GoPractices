package main

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}
type Pair struct {
	Val    int
	VerDis int
}

func main() {
	tree := &Tree{Val: 1, Left: &Tree{Val: 2, Left: &Tree{Val: 4}, Right: &Tree{Val: 5}}, Right: &Tree{Val: 3, Left: &Tree{Val: 6}, Right: &Tree{Val: 7}}}
	displayTopViewOfBinaryTree(tree)
}

func displayTopViewOfBinaryTree(root *Tree) {
	dp := make(map[int]Pair)
	helper(root, dp, 0, 0)
	for _, v := range dp {
		fmt.Print(v.Val, " ")
	}
}
func helper(root *Tree, dp map[int]Pair, d, l int) {
	if root == nil {
		return
	}
	pair, ok := dp[d]
	if !ok {
		dp[d] = Pair{Val: root.Val, VerDis: l}
	} else if pair.VerDis > l {
		dp[d] = Pair{Val: root.Val, VerDis: l}
	}
	helper(root.Left, dp, d-1, l+1)
	helper(root.Right, dp, d+1, l+1)
}
