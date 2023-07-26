package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].priority > p[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (p *PriorityQueue) Push(x any) {
	item := x.(*Item)
	size := len(*p)
	item.index = size
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() any {
	pq := *p
	size := len(pq)
	item := pq[size-1]
	item.index = -1
	pq[size-1] = nil
	*p = pq[0 : size-1]
	return item
}

func main() {
	size := 10
	pq := make(PriorityQueue, size)
	for i := 0; i < size; i++ {
		pq[i] = &Item{
			value:    fmt.Sprintf("%d", i),
			priority: i,
			index:    i,
		}
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Println(" Heap", item.priority, item.value)
	}
}
