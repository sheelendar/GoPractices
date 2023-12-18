package main

import "fmt"

/*
Smallest Substring of All Characters
Given an array of unique characters arr and a string str, Implement a function getShortestUniqueSubstring
that finds the smallest substring of str containing all the characters in arr. Return "" (empty string) if such a substring doesn’t exist.
Come up with an asymptotically optimal solution and analyze the time and space complexities.
Example:
input:  arr = ['x','y','z'], str = "xyyzyzyx"
output: "zyx"
*/
// https://www.geeksforgeeks.org/smallest-window-contains-characters-string/
//https://www.pramp.com/challenge/wqNo9joKG6IJm67B6z34

func main() {
	arr := []string{"x", "y", "z"}
	str := "xyyzyzyx"
	fmt.Print(GetShortestUniqueSubstring(arr, str))
}
func GetShortestUniqueSubstring(arr []string, str string) string {
	size := len(arr)
	// base condition
	if size <= 0 || len(str) <= 0 {
		return ""
	}
	//  put all required characters into a map
	requiredChar := make(map[string]int)
	for i := 0; i < size; i++ {
		requiredChar[arr[i]]++
	}
	// find the length of required character.
	requiredLength := len(requiredChar)
	count := 0
	start := 0
	result := ""
	dp := make(map[byte]int)
	for end := 0; end < len(str); end++ {
		// check if new character exist into  requiredChar map then add it to dp and increase first time count++;
		if _, ok := requiredChar[string(str[end])]; ok {
			if count < requiredLength && dp[str[end]] == 0 {
				count++
			}
			dp[str[end]]++
		}
		// if count is equal to requiredLength then start excluding character from beginning
		for count == requiredLength {
			// check if result is empty or current length is small from pre result
			if result == "" || (end-start) < len(result) {
				result = str[start : end+1]
			}
			// update dp and counter
			if dp[str[start]] > 1 {
				dp[str[start]]--
			} else if dp[str[start]] == 1 {
				dp[str[start]]--
				count--
			}
			start++
		}
	}
	return result
}
