package main

import (
	"fmt"
)

type Node struct {
	Val  int
	next *Node
}

// 1->3->7->10->19, k=2, 10
func main() {
	list := Node{Val: 1, next: &Node{Val: 3, next: &Node{Val: 7, next: &Node{Val: 10, next: &Node{Val: 19}}}}}
	k := 5
	fmt.Println(findNElement(&list, k))
}
func findNElement(list *Node, k int) int {
	if list == nil || k == 0 {
		return -1
	}
	pre := list
	next := list
	count := k
	for count > 0 && next != nil {
		next = next.next
		count--
	}
	if next == nil {
		if count == 0 {
			return list.Val
		}
		return -1
	}
	for next != nil {
		pre = pre.next
		next = next.next
	}
	return pre.Val
}
