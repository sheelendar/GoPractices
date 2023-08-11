package main

import "fmt"

func main() {
	node := &Node{Val: 1, Neighbors: []*Node{&Node{Val: 2}, &Node{Val: 3}, &Node{Val: 4}}}

	n := cloneGraph(node)
	displayNode(node)
	displayNode(n)

	fmt.Println()
	fmt.Println()
	n = cloneGraph1(node)
	displayNode(n)
}

func displayNode(n *Node) {
	fmt.Println(n.Val)
	for i := 0; i < len(n.Neighbors); i++ {
		fmt.Print(n.Neighbors[i].Val)
		fmt.Print(" ")
	}
	fmt.Println()
}

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	return iterativeDfs(node)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	dp := make(map[*Node]*Node)
	return dfs(node, dp)
}

func dfs(node *Node, dp map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	n := Node{Val: node.Val}
	dp[node] = &n

	for i := 0; i < len(node.Neighbors); i++ {
		if _, ok := dp[node.Neighbors[i]]; !ok {
			dfs(node.Neighbors[i], dp)
		}
		n.Neighbors = append(n.Neighbors, dp[node.Neighbors[i]])
	}
	return &n
}

func iterativeDfs(node *Node) *Node {
	if node == nil {
		return nil
	}
	n := Node{
		Val: node.Val,
	}
	dp := map[*Node]*Node{}
	dp[node] = &n
	var stack []*Node
	stack = append(stack, node)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := 0; i < len(item.Neighbors); i++ {
			temp := item.Neighbors[i]
			stack = append(stack, temp)
			newNode, ok := dp[temp]
			if !ok {
				newNode = &Node{
					Val: temp.Val,
				}
				dp[temp] = newNode
			}
			dp[item].Neighbors = append(dp[item].Neighbors, newNode)
		}
	}
	return &n
}
