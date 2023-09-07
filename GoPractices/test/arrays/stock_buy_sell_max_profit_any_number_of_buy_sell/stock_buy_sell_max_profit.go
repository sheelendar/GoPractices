package main

import "fmt"

/*

The cost of a stock on each day is given in an array. Find the maximum profit that you can make by buying and selling on those days.
If the given array of prices is sorted in decreasing order, then profit cannot be earned at all.
Examples:
    Input: arr[] = {100, 180, 260, 310, 40, 535, 695}
    Output: 865
    Explanation: Buy the stock on day 0 and sell it on day 3 => 310 – 100 = 210
                           Buy the stock on day 4 and sell it on day 6 => 695 – 40 = 655
                           Maximum Profit  = 210 + 655 = 865

    Input: arr[] = {4, 2, 2, 2, 4}
*/

func main() {
	arr := []int{100, 180, 260, 310, 40, 535, 695}
	fmt.Print("max profit of a share: ", maxProfitCalculate(arr, len(arr)))
}

func maxProfitCalculate(arr []int, n int) any {
	maxProfit := 0
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			maxProfit = maxProfit + arr[i] - arr[i-1]
		}
	}
	return maxProfit
}
