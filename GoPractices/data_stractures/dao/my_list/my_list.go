package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	list.PushFront(25)
	list.PushFront(30)
	//list.InsertBefore(12 30)
	list.PushBack(40)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
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
