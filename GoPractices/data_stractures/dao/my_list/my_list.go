package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	list.PushFront(25)
	list.PushFront(30)
	list.PushBack(40)
	node := list.Front()
	fmt.Println(node.Value.(int))
	list.Remove(node)

	node = list.Front()
	fmt.Println(node.Value.(int))
	list.Remove(node)

	node = list.Front()
	fmt.Println(node.Value.(int))
	list.Remove(node)
}
