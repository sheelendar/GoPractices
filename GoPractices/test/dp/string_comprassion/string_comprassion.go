package main

import "fmt"

/*

Run-length encoding is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "aabccc" we replace "aa" by "a2" and replace "ccc" by "c3". Thus the compressed string becomes "a2bc3".

Notice that in this problem, we are not adding '1' after single characters.

Given a string s and an integer k. You need to delete at most k characters from s such that the run-length encoded version of s has minimum length.

Find the minimum length of the run-length encoded version of s after deleting at most k characters.

https://leetcode.com/problems/string-compression-ii/description/

Example 1:

Input: s = "aaabcccd", k = 2
Output: 4
Explanation: Compressing s without deleting anything will give us "a3bc3d" of length 6. Deleting any of the characters 'a' or 'c' would at most decrease the length of the compressed string to 5, for instance delete 2 'a' then we will have s = "abcccd" which compressed is abc3d. Therefore, the optimal way is to delete 'b' and 'd', then the compressed version of s will be "a3c3" of length 4.

Example 2:

Input: s = "aabbaa", k = 2
Output: 2
Explanation: If we delete both 'b' characters, the resulting compressed string would be "a4" of length 2.

Example 3:

Input: s = "aaaaaaaaaaa", k = 0
Output: 3
Explanation: Since k is zero, we cannot delete anything. The compressed string is "a11" of length 3.

*/

func main() {
	s := "aaabcccd"
	k := 2
	fmt.Println(getLengthOfOptimalCompression(s, k))
}

func getLengthOfOptimalCompression(s string, K int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = n
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for k := 0; k <= K; k++ {
			if k > 0 { // delete s[i]
				dp[i][k] = Min(dp[i][k], dp[i-1][k-1])
			}
			// keep s[i] concate the same, remove the different
			same := 0
			diff := 0
			for j := i; j >= 1; j-- {
				if s[j-1] == s[i-1] {
					same++
				} else {
					diff++
				}

				if diff > k {
					break
				}
				// // baaaaaaccaaaa  k=2, s[i] = a, j will extend to index 1
				dp[i][k] = Min(dp[i][k], dp[j-1][k-diff]+getLen(same))
			}
		}
	}
	return dp[n][K]
}

func getLen(length int) int {
	if length == 1 {
		return 1
	} else if length < 10 {
		return 2
	} else if length < 100 {
		return 3
	}
	return 4
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
