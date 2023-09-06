package main

import "fmt"

/*
Given an n x n square matrix, find sum of all sub-squares of size k x k where k is smaller than or equal to n.
Examples :
Input:
n = 5, k = 3
arr[][] = { {1, 1, 1, 1, 1},
	   {2, 2, 2, 2, 2},
	   {3, 3, 3, 3, 3},
	   {4, 4, 4, 4, 4},
	   {5, 5, 5, 5, 5},
	};
Output:

	18  18  18
	27  27  27
	36  36  36
*/

func main() {
	n := 5
	k := 3
	arr := [][]int{{1, 1, 1, 1, 1},

		{2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5},
	}
	printAllSumofSubMetrix(arr, n, k)
}

func printAllSumofSubMetrix(arr [][]int, n int, k int) {
	if n < k {
		return
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < k; i++ {
			sum = sum + arr[i][j]
		}
		arr[0][j] = sum
		for i := 1; i < n-k+1; i++ {
			sum = sum + arr[i+k-1][j] - arr[i-1][j]
			res[i][j] = sum
		}
	}

	for i := 0; i < n-k+1; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum = sum + res[i][j]
		}
		print(sum)
		print(" ")
		for j := 1; j < n-k+1; j++ {
			sum = res[i][j+k-1] - res[i][j-1]
			fmt.Print(sum)
			fmt.Print(" ")
		}
		println()
	}
	//displayMatrix(res, n)
}

func displayMatrix(res [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(res[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
