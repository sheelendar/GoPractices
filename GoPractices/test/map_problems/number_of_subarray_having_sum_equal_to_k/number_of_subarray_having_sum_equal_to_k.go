package main

import "fmt"

/*
Given an unsorted array of integers, find the number of subarrays having a sum exactly equal to a given number k.
Examples: (https://www.geeksforgeeks.org/number-subarrays-sum-exactly-equal-k/)

	Input : arr[] = {10, 2, -2, -20, 10}, k = -10
	Output : 3
	Explanation: Subarrays: arr[0…3], arr[1…4], arr[3..4] have a sum exactly equal to -10.
*/

/*
totalSum:=10
sum=10
map[20]=1
totalsum:=12
diff=12-(-10)=22
map[12]:=1
totalsum:=12
map[12]:=1
*/

func main() {
	arr := []int{10, 2, -2, -20, 10}
	k := -10
	fmt.Print(numberOfSubArray(arr, len(arr), k))

}

func numberOfSubArray(arr []int, size int, k int) int {
	count := 0
	totalSum := 0
	dp := make(map[int]int)
	// if diff is 0 then also sum find.
	dp[0]++
	for i := 0; i < size; i++ {
		totalSum = totalSum + arr[i]
		diff := totalSum - k
		if _, ok := dp[diff]; ok {
			count = count + dp[diff]
		}
		dp[totalSum]++
	}
	return count
}
