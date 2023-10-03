package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
}
type PriorityQueue []Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {

	return piq[i].priority > piq[j].priority
}
func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
}
func (piq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*piq = append(*piq, item)
}
func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}
func main() {
	piq := make(PriorityQueue, 0)

	piq.Push(Item{value: "Item 1", priority: 30})
	piq.Push(Item{value: "Item 2", priority: 10})
	piq.Push(Item{value: "Item 3", priority: 20})
	heap.Init(&piq)
	heap.Push(&piq, Item{value: "Item 5", priority: 50})
	heap.Push(&piq, Item{value: "Item 6", priority: 60})
	heap.Push(&piq, Item{value: "Item 4", priority: 0})

	for piq.Len() > 0 {
		item := heap.Pop(&piq).(Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.value, item.priority)
	}
}
