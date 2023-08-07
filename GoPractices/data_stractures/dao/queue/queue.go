package main

import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) Add(data int) {
	q.queue = append(q.queue, data)
}
func (q *Queue) Remove() int {
	if len(q.queue) == 0 {
		return -1
	}
	data := q.queue[0]
	if len(q.queue) == 1 {
		q.queue = nil
		return data
	} else {
		q.queue = q.queue[1:]
	}
	return data
}
func (q *Queue) Peek() int {
	if len(q.queue) == 0 {
		return -1
	}
	return q.queue[0]
}

func (q *Queue) IsEmpty() bool {
	if len(q.queue) == 0 {
		return true
	}
	return false
}

func main() {
	q := Queue{}
	q.Add(10)
	q.Add(12)
	q.Add(15)
	display(q)
	fmt.Println()
	fmt.Println(q.Peek())
	fmt.Println(q.Remove())
	display(q)
}
func display(q Queue) {
	for i := 0; i < len(q.queue); i++ {
		fmt.Print(q.queue[i])
	}
}
