package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/problems/ipo/?envType=study-plan-v2&envId=top-interview-150
Suppose LeetCode will start its IPO soon. In order to sell a good price of its shares to Venture Capital, LeetCode would like to work on some projects to increase its capital before the IPO. Since it has limited resources, it can only finish at most k distinct projects before the IPO. Help LeetCode design the best way to maximize its total capital after finishing at most k distinct projects.
You are given n projects where the ith project has a pure profit profits[i] and a minimum capital of capital[i] is needed to start it.
Initially, you have w capital. When you finish a project, you will obtain its pure profit and the profit will be added to your total capital.
Pick a list of at most k distinct projects from given projects to maximize your final capital, and return the final maximized capital.
The answer is guaranteed to fit in a 32-bit signed integer.

Example 1:
Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
Output: 4
Explanation: Since your initial capital is 0, you can only start the project indexed 0.
After finishing it you will obtain profit 1 and your capital becomes 1.
With capital 1, you can either start the project indexed 1 or the project indexed 2.
Since you can choose at most 2 projects, you need to finish the project indexed 2 to get the maximum capital.
Therefore, output the final maximized capital, which is 0 + 1 + 3 = 4.
*/

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	fmt.Println(findMaximizedCapital(k, w, profits, capital))
}

type Item struct {
	p int
	c int
}
type Pair struct {
	max int
	w   int
	p   int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	if k > n {
		k = n
	}
	var items []Item
	for i := 0; i < n; i++ {
		items = append(items, Item{p: profits[i], c: capital[i]})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].c == items[j].c {
			return items[i].p < items[j].p
		}
		return items[i].c < items[j].c
	})
	//  maxProfit:=w
	var queue []Pair
	i := 0
	queue = append(queue, Pair{max: 0, w: w, p: 0})
	for len(queue) > 0 && k > 0 {
		for i < n && w >= items[i].c {
			queue = append(queue, Pair{p: items[i].p, w: items[i].c})
			i++
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].p > queue[j].p
		})
		w += queue[0].p
		queue = queue[1:]
		k--
	}
	return w
}
