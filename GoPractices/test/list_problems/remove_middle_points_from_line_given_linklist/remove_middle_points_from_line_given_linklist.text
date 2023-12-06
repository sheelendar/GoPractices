package main

import "fmt"

/*
Given a linked list of coordinates where adjacent points either form a vertical line or a horizontal line.
Delete points from the linked list which are in the middle of a horizontal or vertical line.
(https://www.geeksforgeeks.org/given-linked-list-line-segments-remove-middle-points/)
Examples:
Input:  (0,10)->(1,10)->(5,10)->(7,10)

	  |
	(7,5)->(20,5)->(40,5)

Output: Linked List should be changed to following

	(0,10)->(7,10)
	          |
	        (7,5)->(40,5)
*/

type Node struct {
	x    int
	y    int
	next *Node
}

func main() {
	node := &Node{x: 0, y: 10}
	node.next = &Node{x: 1, y: 10}
	node.next.next = &Node{x: 3, y: 10}
	node.next.next.next = &Node{x: 10, y: 10}
	node.next.next.next.next = &Node{x: 10, y: 8}
	node.next.next.next.next.next = &Node{x: 10, y: 5}
	node.next.next.next.next.next.next = &Node{x: 20, y: 5}
	node.next.next.next.next.next.next.next = &Node{x: 40, y: 5}
	displayList(node)
	deleteMiddleNodes(node)
	fmt.Println()
	displayList(node)
}

func displayList(node *Node) {
	if node == nil {
		return
	}
	fmt.Print("x= ", node.x)
	fmt.Print(" ")
	fmt.Print("y= ", node.y)
	fmt.Print("   ")
	displayList(node.next)
}

func deleteMiddleNodes(node *Node) {
	if node == nil || node.next == nil || node.next.next == nil {
		return
	}
	next := node.next
	nextNext := node.next.next

	if node.x == next.x {
		for nextNext != nil && next.x == nextNext.x {
			node.next = next.next
			next.next = nil
			next = nextNext
			nextNext = next.next
		}
	} else if node.y == next.y {
		for nextNext != nil && next.y == nextNext.y {
			node.next = next.next
			next.next = nil
			next = nextNext
			nextNext = next.next
		}
	}
	deleteMiddleNodes(node.next)
}
