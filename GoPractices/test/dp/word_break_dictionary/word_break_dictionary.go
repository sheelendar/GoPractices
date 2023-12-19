package main

import "fmt"

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence
of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/
func main() {
	/*s := "leetcode"
	wordDict := []string{"leet", "code"}*/

	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
	fmt.Println(wordBreakWithDPApproach(s, wordDict))
}

func wordBreakWithDPApproach(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	dic := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		dic[wordDict[i]] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			str := s[j:i]
			if dp[j] && dic[str] {
				dp[i] = true
			}
		}
	}
	return dp[n]
}

func wordBreak(s string, wordDict []string) bool {

	dic := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		dic[wordDict[i]] = true
	}
	return wordUtil(s, dic)
}

// recurrsion solution taking too much time. not good
func wordUtil(s string, dic map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	if len(dic) == 0 {
		return false
	}
	if dic[s] {
		return true
	}
	res := false
	size := len(s)
	for i := 1; i <= size; i++ {
		res = dic[s[0:i]] && wordUtil(s[i:size], dic)
		if res {
			return true
		}
	}
	return res
}
