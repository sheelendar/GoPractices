package main

import "fmt"

/*
Given an array of numbers, return true if given array can represent preorder traversal of a Binary Search Tree,
else return false. Expected time complexity is O(n).
Examples:(https://www.geeksforgeeks.org/check-if-a-given-array-can-represent-preorder-traversal-of-binary-search-tree/)
Input:  pre[] = {2, 4, 3}
Output: true
Input:  pre[] = {2, 4, 1}
Output: false
*/

func main() {
	pre1 := []int64{40, 30, 35, 80, 100}     // output: true
	pre2 := []int64{40, 30, 35, 20, 80, 100} // output: false
	n := len(pre1)
	if canRepresentBST(pre1, n) == true {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}
	if canRepresentBST(pre2, n) == true {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}
}

func canRepresentBST(pre []int64, n int) bool {
	var stack []int64
	stack = append(stack, pre[0])
	i := 1
	old := int64(-9999)
	for ; i < n; i++ {
		// check there is no element less than old in array if there then return false.
		if old > pre[i] {
			return false
		}
		// check if you get more than stack's top value element then pop and replace old element
		for len(stack) > 0 && stack[len(stack)-1] < pre[i] {
			old = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pre[i])
	}
	return true
}
