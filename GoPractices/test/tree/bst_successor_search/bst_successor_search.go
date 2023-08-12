package main

import (
	"fmt"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func FindInOrderSuccessor(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	for node != nil {
		if node.value == key {
			break
		} else if node.value < key {
			node = node.right
		} else {
			node = node.left
		}
	}
	if node == nil {
		return nil
	}
	if node.right != nil {
		node = node.right
		for node.left != nil {
			node = node.left
		}
		return node
	} else {
		parant := node.parent
		for parant.left != node {
			node = parant
			parant = node.parent
		}
		return parant
	}
	return nil
}

func main() {
	root := &Node{value: 20}

	left1 := &Node{value: 8, parent: root}
	root.left = left1
	right1 := &Node{value: 25, parent: root}
	root.right = right1

	left2 := &Node{value: 4, parent: left1}
	left1.left = left2
	right2 := &Node{value: 12, parent: left1}
	left1.right = right2

	left3 := &Node{value: 10, parent: right2}

	right2.left = left3

	right3 := &Node{value: 14, parent: right2}
	right2.right = right3

	n := FindInOrderSuccessor(root, 14)

	if n != nil {
		fmt.Println(n.value)
	}

	n = FindInOrderSuccessor(root, 10)

	if n != nil {
		fmt.Println(n.value)
	}
}
