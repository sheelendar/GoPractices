package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 5, 3, 2, 8, 10, 9}
	K := 3

	//arr := []int{6, 5}
	//K := 3

	sotArray(arr, len(arr), K)
}

func sotArray(arr []int, n int, k int) {
	if n == 0 {
		return
	}
	k = k + 1
	if k > n {
		k = n
	}
	queue := make([]int, k)

	for i := 0; i < k; i++ {
		queue[i] = arr[i]
	}
	sort.Ints(queue)
	i := 0
	for ; i < n-k; i++ {
		arr[i] = queue[0]
		queue = queue[1:]
		queue = append(queue, arr[i+k])
		sort.Ints(queue)
	}
	size := len(queue)
	for k := 0; k < size; k++ {
		arr[i] = queue[0]
		queue = queue[1:]
		i++
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}
