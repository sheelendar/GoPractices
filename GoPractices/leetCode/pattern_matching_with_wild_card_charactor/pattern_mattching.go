package main

import "fmt"

func main() {
	text := "baaabab"
	pattern := "*****ba*****ab"
	dp := make(map[string]bool)
	fmt.Println(IsPatternMatched(text, pattern, len(text)-1, len(pattern)-1, dp))
	fmt.Print(IsMatch(text, pattern))
}

func IsPatternMatched(text string, pattern string, m, n int, dp map[string]bool) bool {
	if m == 0 || n == 0 {
		if n != 0 {
			dp[fmt.Sprintf("%d%d", m, n)] = false
			return false
		}
		dp[fmt.Sprintf("%d%d", m, n)] = true
		return true
	}
	if v, ok := dp[fmt.Sprintf("%d%d", m, n)]; ok {
		return v
	}
	res := false
	if pattern[n] == '.' || pattern[n] == text[m] {
		res := IsPatternMatched(text, pattern, m-1, n-1, dp)
		dp[fmt.Sprintf("%d%d", m, n)] = res
		return res
	} else if pattern[n] == '*' {
		res := IsPatternMatched(text, pattern, m-1, n, dp) || IsPatternMatched(text, pattern, m, n-1, dp)
		dp[fmt.Sprintf("%d%d", m, n)] = res
		return res
	}
	dp[fmt.Sprintf("%d%d", m, n)] = res
	return res
}
func IsPatternMached() {

}

func IsMatch(text string, pattern string) bool {

	N := len(text)
	M := len(pattern)
	if M == 0 && N == 0 {
		return true
	}

	if M == 0 && N != 0 {
		return false
	}
	result := make([][]bool, N+1)

	for i := 0; i <= N; i++ {
		result[i] = make([]bool, M+1)
	}
	result[0][0] = true

	for i := 1; i <= M; i++ {
		result[0][i] = result[0][i-1]
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
