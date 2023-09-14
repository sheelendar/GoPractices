package main

import "fmt"

/*
Given a binary array and an integer m, find the position of zeroes flipping which creates maximum number of consecutive 1’s in array.
Examples :
Input:   arr[] = {1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1}
         m = 2
Output:  5 7
We are allowed to flip maximum 2 zeroes. If we flip
arr[5] and arr[7], we get 8 consecutive 1's which is
maximum possible under given constraints

Input:   arr[] = {1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1}
         m = 1
Output:  7
We are allowed to flip maximum 1 zero. If we flip
arr[7], we get 5 consecutive 1's which is maximum
possible under given constraints.
*/

func main() {
	arr := []int{1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1}
	m := 2
	consecutiveOneMaxmized(arr, m)
}

func consecutiveOneMaxmized(arr []int, m int) {
	n := len(arr)
	left := 0
	right := 0
	bestL := 0
	maxOne := 0
	maxZeros := 0

	for right < n {
		// if maxZeros is less than m then increase maxZeros if you get zero and increase right.
		if maxZeros <= m {
			if arr[right] == 0 {
				maxZeros++
			}
			right++
		}
		// now you have max of all one after flipped given number of zeros.
		// check remaining array
		// if maxZeros is more than m then skip current zeros and increase left.
		if maxZeros > m {
			if arr[left] == 0 {
				maxZeros--
			}
			left++
		}
		// calculate maxOne size if maxZeros <= m is true.
		if maxOne <= right-left && (maxZeros <= m) {
			maxOne = right - left
			bestL = left
		}
	}
	for i := 0; i < maxOne; i++ {
		if arr[bestL+i] == 0 {
			fmt.Print(bestL + i)
			fmt.Print(" ")
		}

	}
}
