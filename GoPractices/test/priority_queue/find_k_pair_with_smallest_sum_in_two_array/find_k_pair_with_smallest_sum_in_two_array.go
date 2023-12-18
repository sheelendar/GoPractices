package main

import (
	"container/heap"
	"fmt"
)

/*
You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.
Define a pair (u, v) which consists of one element from the first array and one element from the second array.
Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.
Example 1: https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/?envType=study-plan-v2&envId=top-interview-150
Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3,    Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
Example 2:
Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2,   Output: [[1,1],[1,1]]
Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
*/
type Item struct {
	i        int
	j        int
	priority int
}
type PriorityQueue []Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {

	return piq[i].priority < piq[j].priority
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
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}

	//nums1 := []int{1, 2}
	//nums2 := []int{3}
	k := 3
	result := kSmallestPairs(nums1, nums2, k)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i][0], "  ", result[i][1])
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	m := len(nums1)
	n := len(nums2)
	if m == 0 || n == 0 {
		return result
	}
	if k > m*n {
		k = m * n
	}
	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	for i := 0; i < m; i++ {
		heap.Push(&queue, Item{priority: nums1[i] + nums2[0], i: i, j: 0})
	}
	fmt.Println(queue)
	for i := 0; i < k; i++ {
		item := heap.Pop(&queue).(Item)
		result = append(result, []int{nums1[item.i], nums2[item.j]})
		nextJ := item.j + 1
		if nextJ < n {
			heap.Push(&queue, Item{priority: nums1[item.i] + nums2[nextJ], i: item.i, j: nextJ})
		}
	}
	return result
}
