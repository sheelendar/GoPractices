package main

import "fmt"

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

	'.' Matches any single character.​​​​
	'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
*/
func main() {
	fmt.Println(isMatch("aa", "a*"))
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] { // when charactor is equal or pattern has dot(.) then check perivios diagonal
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // pa[j-1]=="*" then there two cases
				dp[i][j] = dp[i][j-2]                  // when * count 0 times then remove check 2 character back.
				if s[i-1] == p[j-2] || p[j-2] == '.' { //
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

			}
		}
	}
	return dp[m][n]
}
