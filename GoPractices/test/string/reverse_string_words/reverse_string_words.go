package main

import (
	"fmt"
	"strings"
)

/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/
func main() {
	s := "  hello world  "
	fmt.Println(strings.Trim(s, " "))
	fmt.Println(reverseWords(s))
	var b strings.Builder
	b.WriteString("")
	strings.HasPrefix()
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	fmt.Println(words)
	reverse(words)
	fmt.Println(words)
	return strings.Join(words, " ")
}

func reverse(arr []string) {
	if len(arr) == 0 {
		return
	}
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
