package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := &ListNode{Val: 1}
	removeNthFromEnd(head, 1)
	curr := head
	for curr != nil {
		fmt.Print(" ", curr.Val)
		curr = curr.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	count := n
	curr := head
	for count > 0 && curr != nil {
		count--
		curr = curr.Next
	}
	if curr == nil {
		head = head.Next
		return head
	}
	temp := head
	var pre *ListNode
	for curr != nil {
		curr = curr.Next
		pre.Next = temp
		temp = temp.Next
	}
	pre.Next = temp.Next
	return head
}
