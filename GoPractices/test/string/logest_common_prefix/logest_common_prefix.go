package main

import (
	"fmt"
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	small := strs[0]
	for i := 1; i < n; i++ {
		if len(small) > len(strs[i]) {
			small = strs[i]
		}
	}

	for i := len(small); i > 0; i-- {
		str := small[0:i]
		j := 0
		for ; j < n; j++ {
			if !strings.HasPrefix(strs[j], str) {
				break
			}
		}
		if j == n {
			return str
		}
	}
	return ""
}
