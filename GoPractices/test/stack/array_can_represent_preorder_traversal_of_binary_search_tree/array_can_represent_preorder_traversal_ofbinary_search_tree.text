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

type Stack struct {
	stack []int64
}

func (s *Stack) Push(item int64) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() int64 {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return -1
}
func (s *Stack) Top() int64 {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return -1
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}

func main() {
	pre1 := []int64{40, 30, 35, 80, 100}
	pre2 := []int64{40, 30, 35, 20, 80, 100}
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
	stack := Stack{}
	stack.Push(pre[0])
	i := 1
	old := int64(-9999)
	for ; i < n; i++ {
		// check if you get more than stack's top value element then pop and replace old element
		for !stack.IsEmpty() && stack.Top() < pre[i] {
			old = stack.Pop()
		}
		stack.Push(pre[i])

		// check there is no element less than old in array if there then return false.
		if old > pre[i] {
			return false
		}
	}

	return true

}
