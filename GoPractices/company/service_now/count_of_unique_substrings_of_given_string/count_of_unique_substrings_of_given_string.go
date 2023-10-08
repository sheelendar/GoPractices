/*
Count number of Distinct Substring in a String
Read
Discuss
Courses
Practice

Given a string, count all distinct substrings of the given string.

Examples:

Input : abcd
Output : abcd abc ab a bcd bc b cd c d  // 10
All Elements are Distinct

Input : aaa
Output : aaa aa a aa a a   /// 3
All elements are not Distinct

Input: kincenvizh
output: 53

https://www.geeksforgeeks.org/count-number-of-distinct-substring-in-a-string/
*/
package main

import "fmt"

func main() {
	fmt.Println(countUniqueSubstrings("aaa"))
	fmt.Println(countUniqueSubstrings("kincenvizh"))
	fmt.Println(countUniqueSubstrings("abcd"))
}

func countUniqueSubstrings(s string) int {
	dp := make(map[string]bool)
	arr := []byte(s)
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			a := string(arr[i:j])
			if _, ok := dp[a]; !ok {
				dp[a] = true
			}
		}
	}
	return len(dp)
}
