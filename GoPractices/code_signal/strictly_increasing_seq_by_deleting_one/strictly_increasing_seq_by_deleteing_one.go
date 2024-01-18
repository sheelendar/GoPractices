package main

import "fmt"

/*
Given a sequence of integers as an array, determine whether it is possible to obtain a strictly
increasing sequence by removing no more than one element from the array.
Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an.
Sequence containing only one element is also considered to be strictly increasing.

Example

	    For sequence = [1, 3, 2, 1], the output should be
	    solution(sequence) = false.

	    There is no one element in this array that can be removed in order to get a strictly increasing sequence.
	    For sequence = [1, 3, 2], the output should be
	    solution(sequence) = true.
	    You can remove 3 from the array to get the strictly increasing sequence [1, 2]. Alternately,
		you can remove 2 to get the strictly increasing sequence [1, 3].
*/
func main() {
	arr := []int{1, 3, 2, 1}
	fmt.Println(solution(arr))
}

func solution(sequence []int) bool {
	n := len(sequence)
	count := 0
	index := -1
	for i := 1; i < n; i++ {
		if sequence[i-1] >= sequence[i] {
			count++
			index = i
		}
	}
	if count == 0 {
		return true
	}
	if count > 1 {
		return false
	}
	if index == 1 || index == n-1 {
		return true
	}
	if sequence[index-1] < sequence[index+1] {
		return true
	}
	if index-2 >= 0 && sequence[index-2] < sequence[index] {
		return true
	}
	return false
}
