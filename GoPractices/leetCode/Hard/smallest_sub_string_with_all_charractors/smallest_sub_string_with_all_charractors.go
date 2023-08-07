package main

import "fmt"

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
	requiredChar := make(map[string]bool)
	for i := 0; i < size; i++ {
		requiredChar[arr[i]] = true
	}
	// find the length of required character.
	requiredLength := len(requiredChar)
	count := 0
	start := 0
	result := ""
	dp := make(map[byte]int)
	for end := 0; end < len(str); end++ {
		// check if new character exist into  requiredChar map then add it to dp and increase first time count++;
		if requiredChar[string(str[end])] {
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
