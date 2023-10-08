package main

import (
	"fmt"
	"strings"
)

func main() {
	var result [][]string
	solution("aab", "", &result)
	fmt.Println(result)
}

func solution(str, asf string, result *[][]string) {
	if len(str) == 0 {
		arr := strings.Split(asf, ",")
		*result = append(*result, arr[1:])
		fmt.Println(asf)
		return
	}
	for i := 0; i < len(str); i++ {
		prefix := str[0 : i+1]
		ros := str[i+1:]
		if palindrome(prefix) {
			solution(ros, asf+","+prefix, result)
		}
	}
}

func palindrome(str string) bool {
	i := 0
	j := len(str) - 1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
