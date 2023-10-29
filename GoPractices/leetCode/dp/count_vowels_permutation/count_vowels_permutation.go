package main

import "fmt"

/*
https://leetcode.com/problems/count-vowels-permutation/
Given an integer n, your task is to count how many strings of length n can be formed under the following rules:

	Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
	Each vowel 'a' may only be followed by an 'e'.
	Each vowel 'e' may only be followed by an 'a' or an 'i'.
	Each vowel 'i' may not be followed by another 'i'.
	Each vowel 'o' may only be followed by an 'i' or a 'u'.
	Each vowel 'u' may only be followed by an 'a'.

Since the answer may be too large, return it modulo 10^9 + 7.
Example 1:
Input: n = 1
Output: 5
Explanation: All possible strings are: "a", "e", "i" , "o" and "u".

Example 2:
Input: n = 2
Output: 10
Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".

Example 3:
Input: n = 5
Output: 68
*/
func main() {
	fmt.Println(countVowelPermutation(2))
}
func countVowelPermutation(n int) int {
	/*dp := make(map[byte][]byte)
	dp['a'] = []byte{'e'}
	dp['e'] = []byte{'a', 'i'}
	dp['i'] = []byte{'a', 'e', 'o', 'u'}
	dp['o'] = []byte{'i', 'u'}
	dp['u'] = []byte{'a'}*/
	return helper(n)
}

func helper(n int) int {
	mod := 1000000007
	a := 1
	e := 1
	i := 1
	o := 1
	u := 1
	for k := 0; k < n-1; k++ {
		a_next := e
		e_next := (a + i) % mod
		i_next := (a + e + o + u) % mod
		o_next := (i + u) % mod
		u_next := a
		a, e, i, o, u = a_next, e_next, i_next, o_next, u_next

	}
	return (a + e + i + o + u) % mod
}

/*
func helper(dp map[byte][]byte, str string, index, n int, count *int) {
	if index == n {
		*count = *count + 1
		return
	}
	if index == 0 {
		for ch, _ := range dp {
			helper(dp, str+string(ch), index+1, n, count)
		}
	} else {
		adj := dp[str[index-1]]
		for i := 0; i < len(adj); i++ {
			helper(dp, str+string(adj[i]), index+1, n, count)
		}
	}
}
*/
