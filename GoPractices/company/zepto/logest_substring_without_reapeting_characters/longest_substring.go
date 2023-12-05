package main

import "fmt"

/*
Given a string s, find the length of the longest substring without repeating characters.
Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func main() {
	s := "abcabcbb"
	//Output: 3
	fmt.Println(findMaxLength(s))
}
func findMaxLength(str string) int {
	ans := 0
	n := len(str)
	index := -1
	dp := make(map[byte]int)

	for i := 0; i < n; i++ {
		if v, ok := dp[str[i]]; ok {
			index = Max(v, index)
		}
		dp[str[i]] = i
		ans = Max(ans, (i - index))
	}
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
