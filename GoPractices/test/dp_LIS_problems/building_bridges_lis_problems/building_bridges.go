package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Consider a 2-D map with a horizontal river passing through its center. There are n cities on the southern bank with x-coordinates
a(1) … a(n) and n cities on the northern bank with x-coordinates b(1) … b(n). You want to connect as many north-south pairs of
cities as possible with bridges such that no two bridges cross. When connecting cities, you can only connect city a(i) on the northern
bank to city b(i) on the southern bank. Maximum number of bridges that can be built to connect north-south pairs with the
above mentioned constraints.
The values in the upper bank can be considered as the northern x-coordinates of the cities and the values in the bottom bank
can be considered as the corresponding southern x-coordinates of the cities to which the northern x-coordinate city can be connected.
Examples:
Input : 6 4 2 1

	2 3 6 5

Output : Maximum number of bridges = 2
Explanation: Let the north-south x-coordinates
be written in increasing order.

1  2  3  4  5  6

	\  \
	 \  \        For the north-south pairs
	  \  \       (2, 6) and (1, 5)
	   \  \      the bridges can be built.
	    \  \     We can consider other pairs also,
	     \  \    but then only one bridge can be built
	      \  \   because more than one bridge built will
	       \  \  then cross each other.
	        \  \

1  2  3  4  5  6

Input : 8 1 4 3 5 2 6 7

	1 2 3 4 5 6 7 8

Output : Maximum number of bridges = 5
*/
func main() {
	//southernBank := []int{8, 1, 4, 3, 5, 2, 6, 7}  // output : 5
	//northernBank := []int{1, 2, 3, 4, 5, 6, 7, 8}

	southernBank := []int{6, 4, 2, 1} // output : 2
	northernBank := []int{2, 3, 6, 5}
	fmt.Println(maxNumberOfBuildingBridge(southernBank, northernBank))
}

type NorthSouth struct {
	North int
	South int
}

func maxNumberOfBuildingBridge(a []int, b []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	var banks []NorthSouth
	for i := 0; i < n; i++ {
		banks = append(banks, NorthSouth{North: a[i], South: b[i]})
	}
	sort.Slice(banks, func(i, j int) bool {
		return banks[i].North < banks[j].North
	})
	dp := make([]int, n)
	dp[0] = 1
	max := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if banks[j].South < banks[i].South && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				max = int(math.Max(float64(max), float64(dp[i])))
			}
		}
	}
	return max
}
