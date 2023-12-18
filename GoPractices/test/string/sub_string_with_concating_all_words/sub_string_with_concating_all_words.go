package main

import "fmt"

/*
You are given a string s and an array of strings words. All the strings of words are of the same length.
A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.

	    For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are
		 all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation of any permutation of words.

		 Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

https://leetcode.com/problems/substring-with-concatenation-of-all-words
Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Since words.length == 2 and words[i].length == 3, the concatenated substring has to be of length 6.
The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.
The output order does not matter. Returning [9,0] is fine too.

Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Explanation: Since words.length == 4 and words[i].length == 4, the concatenated substring has to be of length 16.
There is no substring of length 16 in s that is equal to the concatenation of any permutation of words.
We return an empty array.
*/
func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	n := len(words)
	var res []int
	dp := make(map[string]int)
	for _, w := range words {
		dp[w]++
	}
	m := len(s)
	size := len(words[0])
	for i := 0; i <= m-(n*size); i++ {
		cache := deepCopyOfMap(dp)
		count := 0
		for j := i; j < m; {
			sub := s[j : j+size]
			v, ok := cache[sub]
			if !ok || v == 0 {
				break
			}
			count++
			cache[sub]--
			if count == n {
				res = append(res, i)
				break
			}
			j = j + size
		}
	}
	return res
}

func deepCopyOfMap(dp map[string]int) map[string]int {
	newDP := make(map[string]int)
	for k, v := range dp {
		newDP[k] = v
	}
	return newDP
}
