package main

import "fmt"

/*
You’re testing a new driverless car that is located at the Southwest (bottom-left)
corner of an n×n grid. The car is supposed to get to the opposite, Northeast (top-right),
corner of the grid. Given n, the size of the grid’s axes, write a function numOfPathsToDest
that returns the number of the possible paths the driverless car can take.

For convenience, let’s represent every square in the grid as a pair (i,j). The first coordinate
in the pair denotes the east-to-west axis, and the second coordinate denotes the south-to-north axis.
The initial state of the car is (0,0), and the destination is (n-1,n-1).

The car must abide by the following two rules: it cannot cross the diagonal border.
In other words, in every step the position (i,j) needs to maintain i >= j.
See the illustration above for n = 5. In every step, it may go one square North (up),
or one square East (right), but not both. E.data_stractures. if the car is at (3,1), it may go to (3,2) or (4,1).

Explain the correctness of your function, and analyze its time and space complexities.

Example:

input:  n = 4

output: 5
*/
func main() {
	fmt.Println(NumOfPathsToDest(4))
}

func NumOfPathsToDest(n int) int {
	dp := make(map[string]int)
	return Helper(n-1, n-1, dp)
}

func Helper(m, n int, dp map[string]int) int {
	if m < 0 || n < 0 {
		return 0
	}
	// base condition means that every path reached to it's destination.
	if m == 0 && n == 0 {
		return 1
	}
	if v, ok := dp[fmt.Sprintf("%d_%d", m, n)]; ok {
		return v
	}
	if m > n {
		res := Helper(m, n-1, dp) + Helper(m-1, n, dp)
		dp[fmt.Sprintf("%d_%d", m, n)] = res
		return res
	}
	res := Helper(m, n-1, dp)
	dp[fmt.Sprintf("%d_%d", m, n)] = res
	return res
}
