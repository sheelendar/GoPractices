package main

import (
	"container/list"
	"fmt"
)

func main() {
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Print(lru.Get(4))
	fmt.Print(lru.Get(3))
	fmt.Print(lru.Get(2))
	fmt.Print(lru.Get(1))
	lru.Put(5, 5)
	fmt.Print(lru.Get(1))
	fmt.Print(lru.Get(2))
	fmt.Print(lru.Get(3))
	fmt.Print(lru.Get(4))
	fmt.Print(lru.Get(5))

	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
}

type Node struct {
	// key is only needed to delete entry in data map
	// when removing LRU item
	k int
	v int
}

type LRUCache struct {
	capacity int
	data     map[int]*list.Element
	lLits    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]*list.Element),
		lLits:    list.New(),
	}
}
func (c *LRUCache) Put(key int, value int) {

	if ele, ok := c.data[key]; ok {
		ele.Value = Node{k: key, v: value}
		c.lLits.MoveToFront(ele)
		return
	}
	if c.lLits.Len() == c.capacity {
		last := c.lLits.Back()
		node := last.Value.(Node) // typecast into Node
		delete(c.data, node.k)
		c.lLits.Remove(last)
	}
	c.data[key] = c.lLits.PushFront(Node{k: key, v: value})
}

func (c *LRUCache) Get(key int) int {
	if ele, ok := c.data[key]; ok {
		c.lLits.MoveToFront(ele)
		return ele.Value.(Node).v
	}
	return -1
}
