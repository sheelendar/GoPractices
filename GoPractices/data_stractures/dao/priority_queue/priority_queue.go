package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}
type PriorityQueue struct {
	priorityQueue []*Item
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.priorityQueue[i], pq.priorityQueue[j] = pq.priorityQueue[j], pq.priorityQueue[i]
	pq.priorityQueue[i].index = i
	pq.priorityQueue[j].index = j
}

func (pq PriorityQueue) Len() int { return len(pq.priorityQueue) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.priorityQueue[i].priority > pq.priorityQueue[j].priority
}
func (pq *PriorityQueue) Push(data any) {
	item := data.(*Item)
	n := len(pq.priorityQueue)
	item.index = n
	pq.priorityQueue = append(pq.priorityQueue, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.priorityQueue
	n := len(old)
	var item *Item
	if n == 1 {
		item = old[n-1]
		old[n-1] = nil // avoid memory leak
		item.index = -1
	} else {
		item = old[n-1]
		old[n-1] = nil     // avoid memory leak
		item.index = n - 1 // for safety
		pq.priorityQueue = old[1 : n-1]
	}
	return item
}

func main() {
	// Some items and their priorities.
	items := map[int]int{
		1: 3, 5: 10, 2: 2, 3: 4,
	}
	pq := &PriorityQueue{}
	i := 0
	for value, priority := range items {
		pq.Push(&Item{
			value:    value,
			priority: priority,
			index:    i,
		})
		i++
	}
	heap.Init(pq)
	for i := 0; i < pq.Len(); i++ {
		fmt.Print(pq.priorityQueue[i].priority)
		fmt.Print(" ")
	}
}
