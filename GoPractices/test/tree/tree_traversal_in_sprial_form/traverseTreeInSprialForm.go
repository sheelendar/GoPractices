package main

import "fmt"

/*type Stack struct {
	stack []*Node
}

func (s *Stack) Push(item *Node) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() *Node {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return nil
}
func (s *Stack) Top() *Node {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return nil
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}*/

type Node struct {
	data  int64
	left  *Node
	right *Node
}

func main() {
	node := Node{data: 1}
	node.left = &Node{data: 2}
	node.right = &Node{data: 3}
	node.left.left = &Node{data: 7}
	node.left.right = &Node{data: 6}
	node.right.left = &Node{data: 5}
	node.right.right = &Node{data: 4}
	displayTreeInSprilForm(&node)
}

func displayTreeInSprilForm(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(" ")
	var s1 []*Node
	var s2 []*Node
	s1 = append(s1, n)
	for len(s1) != 0 || len(s2) != 0 {
		for len(s1) != 0 {
			node := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			fmt.Print(node.data)
			fmt.Print(" ")
			if node.left != nil {
				s2 = append(s2, node.right)
			}
			if node.right != nil {
				s2 = append(s2, node.left)
			}
		}
		for len(s2) != 0 {
			node := s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
			fmt.Print(node.data)
			fmt.Print(" ")
			if node.left != nil {
				s1 = append(s1, node.left)
			}
			if node.right != nil {
				s1 = append(s1, node.right)
			}
		}
	}
}
