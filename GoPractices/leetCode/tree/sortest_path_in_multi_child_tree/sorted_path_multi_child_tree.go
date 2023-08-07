package main

import "fmt"

type Node struct {
	cost     int
	parent   *Node
	children []*Node
}

func GetCheapestCost(rootNode *Node) int {

	if rootNode == nil {
		return 0
	}
	return minPath(rootNode)
}

func minPath(node *Node) int {
	if node == nil {
		return 0
	}
	if len(node.children) == 0 {
		return node.cost
	}
	min := 999999
	sum := 0
	for i := 0; i < len(node.children); i++ {
		sum = node.cost + minPath(node.children[i])

		if sum < min {
			min = sum
		}
	}
	return min
}

func main() {
	root := Node{cost: 0}
	root.children = []*Node{{cost: 5}, {cost: 3}, {cost: 6}}

	root.children[0].children = []*Node{{cost: 4}}
	root.children[1].children = []*Node{{cost: 2}, {cost: 0}}
	root.children[2].children = []*Node{{cost: 1}, {cost: 5}}

	root.children[1].children[0].children = []*Node{{cost: 1}}
	root.children[1].children[0].children[0].children = []*Node{{cost: 1}}
	root.children[1].children[1].children = []*Node{{cost: 10}}

	fmt.Print(GetCheapestCost(&root))
}
