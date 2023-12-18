package main

import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
	t      interface{}
}

func FindLargestSmallerKey(rootNode *Node, num int) int {
	if rootNode == nil {
		return -1
	}
	result := -1
	for rootNode != nil {
		if rootNode.key < num {
			result = rootNode.key
			rootNode = rootNode.right
		} else {
			rootNode = rootNode.left
		}

	}
	return result
}

func search(root *Node, num int) int {

	if root.left == nil && root.right == nil {

		if root.key < num {
			return root.key
		} else {
			return -1
		}

	} else if root.left != nil && root.right == nil && root.key > num {

		return search(root.left, num)

	} else if root.right != nil && root.key < num {

		res := search(root.right, num)
		if res == -1 {
			return root.key
		}

	} else if root.left != nil {

		return search(root.left, num)
	}

	return -1

}

func main() {

	node := Node{key: 20}
	node.left = &Node{key: 9}
	node.left.left = &Node{key: 5}
	node.left.right = &Node{key: 12}
	node.left.right.left = &Node{key: 11}
	node.left.right.right = &Node{key: 14}
	node.right = &Node{key: 25}

	fmt.Println(FindLargestSmallerKey(&node, 12))

}
