package main

import "fmt"

type Stack struct {
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
}

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
	//fmt.Print(n.data)
	fmt.Print(" ")
	//preOrderTraverse(n.left)
	//preOrderTraverse(n.right)
	s1 := Stack{}
	s2 := Stack{}
	s1.Push(n)
	for !s1.IsEmpty() || !s2.IsEmpty() {

		for !s1.IsEmpty() {
			n := s1.Pop()
			fmt.Print(n.data)
			fmt.Print(" ")
			if n.left != nil {
				s2.Push(n.right)
			}
			if n.right != nil {
				s2.Push(n.left)
			}
		}
		for !s2.IsEmpty() {
			n := s2.Pop()
			fmt.Print(n.data)
			fmt.Print(" ")
			if n.left != nil {
				s1.Push(n.left)
			}
			if n.right != nil {
				s1.Push(n.right)
			}
		}
	}
}
