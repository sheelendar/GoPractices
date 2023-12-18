package main

import "fmt"

/*
? 	The question mark indicates zero or one occurrences of the preceding element. For example, colou?r matches both "color" and "colour".
* 	The asterisk indicates zero or more occurrences of the preceding element. For example, ab*c matches "ac", "abc", "abbc", "abbbc", and so on.
+ 	The plus sign indicates one or more occurrences of the preceding element. For example, ab+c matches "abc", "abbc", "abbbc", and so on, but not "ac".
.   The wildcard . matches any character. For example,
*/
func main() {
	text := "aa"
	pattern := "a"
	dp := make(map[string]bool)
	fmt.Println(IsPatternMatched(text, pattern, 0, 0, dp))
	fmt.Print(IsMatch(text, pattern))
}

func IsPatternMatched(text string, pattern string, m, n int, dp map[string]bool) bool {
	if m >= len(text) {
		if n >= len(pattern) {
			return true
		} else if n+1 < len(pattern) && pattern[n+1] == '*' {
			res := IsPatternMatched(text, pattern, m, n+2, dp)
			return res
		} else {
			return false
		}
	} else if n >= len(pattern) && m < len(text) {
		return false
	}
	if v, ok := dp[fmt.Sprintf("%d_%d", m, n)]; ok {
		return v
	}
	if n+1 < len(pattern) && pattern[n+1] == '*' {
		if pattern[n] == '.' || pattern[n] == text[m] {
			res := IsPatternMatched(text, pattern, m+1, n, dp) || IsPatternMatched(text, pattern, m, n+2, dp)
			dp[fmt.Sprintf("%d_%d", m, n)] = res
			return res
		} else {
			res := IsPatternMatched(text, pattern, m, n+2, dp)
			dp[fmt.Sprintf("%d_%d", m, n)] = res
			return res
		}
	}
	if pattern[n] == '.' || pattern[n] == text[m] {
		res := IsPatternMatched(text, pattern, m+1, n+1, dp)
		dp[fmt.Sprintf("%d_%d", m, n)] = res
		return res
	}
	dp[fmt.Sprintf("%d_%d", m, n)] = false
	return false
}

func IsMatch(text string, pattern string) bool {
	N := len(text)
	M := len(pattern)
	//  base condition if both reach to end then there is match
	if M == 0 && N == 0 {
		return true
	}
	// if pattern is end reached to end and text is pending then there is no match
	if M == 0 && N != 0 {
		return false
	}
	result := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		result[i] = make([]bool, M+1)
	}
	// base condition if both text and pattern are empty
	result[0][0] = true
	for i := 1; i <= M; i++ {
		if pattern[i-1] == '*' {
			result[0][i] = result[0][i-1]
		}
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if pattern[j-1] == '*' {
				result[i][j] = result[i][j-1] || result[i-1][j]
			} else if pattern[j-1] == '.' || text[i-1] == pattern[j-1] {
				result[i][j] = result[i-1][j-1]
			} else {
				result[i][j] = false
			}
		}
	}
	return result[N][M]
}
