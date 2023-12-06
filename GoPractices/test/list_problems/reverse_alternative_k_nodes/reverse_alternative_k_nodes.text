package main

type Node struct {
	data int64
	next *Node
}

func main() {
	n := Node{data: 10, next: &Node{data: 20, next: &Node{30, &Node{40, &Node{50, &Node{60, &Node{70, nil}}}}}}}
	k := 3
	displayList(&n)
	println()
	node := reverseListByKNodes(&n, k)
	displayList(node)
}

func displayList(curr *Node) {
	for ; curr != nil; curr = curr.next {
		print(curr.data)
		print(" ")
	}
}

func reverseListByKNodes(n *Node, k int) *Node {
	if n == nil {
		return nil
	}
	curr := n
	c := k
	var pre, next *Node
	for curr != nil && c > 0 {
		next = curr.next
		curr.next = pre
		pre = curr
		curr = next
		c--
	}
	if n != nil {
		n.next = curr
	}
	c = k
	for ; curr != nil && c > 0; curr = curr.next {
		c--
	}
	if curr != nil {
		reverseListByKNodes(curr.next, k)
	}
	return pre
}
