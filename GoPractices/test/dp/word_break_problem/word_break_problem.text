package main

import (
	"fmt"
	"strings"
)

func main() {
	temp_dictionary := []string{"mobile", "samsung", "sam", "sung",
		"man", "mango", "icecream", "and",
		"go", "i", "like", "ice", "cream"}
	search := "ilikesamsung"
	fmt.Print(wordBreack(search, temp_dictionary, len(temp_dictionary)))
}

func wordBreack(search string, dictionary []string, n int) bool {
	if len(search) == 0 {
		return true
	}
	if n == 0 {
		return true
	}
	dp := make([]bool, len(search)+1)
	dp[0] = true

	for i := 1; i <= len(search); i++ {

		for j := 0; j < i; j++ {
			s1 := search[j:i]

			for k := 0; k < n; k++ {
				if dp[j] && strings.EqualFold(dictionary[k], s1) {
					dp[i] = true
					break
				}
			}

		}
	}
	return dp[len(search)]
}
