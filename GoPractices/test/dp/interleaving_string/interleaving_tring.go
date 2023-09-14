package main

import "fmt"

/*
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where s and t are divided into n and m
substrings
respectively, such that:

	s = s1 + s2 + ... + sn
	t = t1 + t2 + ... + tm
	|n - m| <= 1
	The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...

Note: a + b is the concatenation of strings a and b.

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Explanation: One way to obtain s3 is:
Split s1 into s1 = "aa" + "bc" + "c", and s2 into s2 = "dbbc" + "a".
Interleaving the two splits, we get "aa" + "dbbc" + "bc" + "a" + "c" = "aadbbcbcac".
Since s3 can be obtained by interleaving s1 and s2, we return true.

Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
Explanation: Notice how it is impossible to interleave s2 with any other string to obtain s3.
*/
func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	isInterleave(s1, s2, s3)
}
func isInterleave(s1 string, s2 string, s3 string) {
	n1 := len(s1)
	n2 := len(s2)
	n3 := len(s3)
	fmt.Println(dpHelper(s1, s2, s3, n1, n2, n3))
	dp := make(map[string]bool)
	fmt.Println(helper(s1, s2, s3, 0, 0, 0, n1, n2, n3, dp))
}

// DP Solution in go
func dpHelper(s1, s2, s3 string, n1, n2, n3 int) bool {
	if n1+n2 != n3 {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				if s2[j-1] == s3[i+j-1] {
					dp[i][j] = dp[i][j-1]
				}
			} else if j == 0 {
				if s1[i-1] == s3[i+j-1] {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				dp[i][j] = (s1[i-1] == s3[i+j-1] && dp[i-1][j]) ||
					(s2[j-1] == s3[i+j-1] && dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

// recurssion solution in go

func helper(s1, s2, s3 string, i, j, k, n1, n2, n3 int, dp map[string]bool) bool {
	if i == n1 && j == n2 && k == n3 {
		return true
	}
	key := fmt.Sprintf("%d_%d", i, j)
	if v, ok := dp[key]; ok {
		return v
	}
	if k == n3 {
		return false
	}
	res := false
	if i < n1 && s1[i] == s3[k] {
		res = helper(s1, s2, s3, i+1, j, k+1, n1, n2, n3, dp)
	}
	if j < n2 && s2[j] == s3[k] {
		res = res || helper(s1, s2, s3, i, j+1, k+1, n1, n2, n3, dp)
	}
	dp[key] = res
	return res
}
