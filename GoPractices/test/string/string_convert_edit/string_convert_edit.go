package main

import "fmt"

func main() {

	source := "ABCDEFG"
	target := "ABDFFGH"
	m := len(source)
	n := len(target)
	dp := make([][]int, m+1, n+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	checkEditDistence(m, n, source, target, dp)
	res := diffBetweenTwoStrings(source, target, dp)
	fmt.Println(res)
}

func checkEditDistence(m, n int, str1, str2 string, dp [][]int) {
	if n == 0 && m == 0 {
		return
	}
	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {

			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func diffBetweenTwoStrings(source, target string, dp [][]int) []string {
	var ans []string
	i := 0
	j := 0

	for i < len(source) && j < len(target) {
		if source[i] == target[j] {
			ans = append(ans, string(source[i]))
			i = i + 1
			j = j + 1
		} else {

			if dp[i+1][j] <= dp[i][j+1] {
				ans = append(ans, fmt.Sprintf("-%s", string(source[i])))
				i = i + 1
			} else {
				ans = append(ans, fmt.Sprintf("+%s", string(target[j])))
				j = j + 1
			}
		}
	}
	for j < len(target) {
		ans = append(ans, fmt.Sprintf("+%s", string(target[j])))
		j += 1
	}
	return ans
}

//["A", "B", "-C", "D", "-E", "F", "+F", "G", "+H"
