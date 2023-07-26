package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	firstIndex := 0
	tree := Node{data: 10, left: &Node{
		data: 2, left: &Node{
			data: 7,
		}, right: &Node{
			data: 8,
		},
	}, right: &Node{
		data: 3,
		left: &Node{
			data: 12,
		},
		right: &Node{
			data: 15, left: &Node{data: 14},
		},
	},
	}
	displayLeftView(&tree, &firstIndex, 1)
}

func displayLeftView(tree *Node, firstIndex *int, index int) {
	if tree == nil {
		return
	}
	if *firstIndex < index {
		fmt.Print(tree.data)
		fmt.Print(" ")
		*firstIndex = index
	}
	displayLeftView(tree.left, firstIndex, index+1)
	displayLeftView(tree.right, firstIndex, index+1)

}
