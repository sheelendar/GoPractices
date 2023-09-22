package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"rat", "abcd", "cat"}
	result := solution(a)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func solution(a []string) []string {
	var result []string
	for i := 0; i < len(a); i++ {
		generateAllAnagrams(a[i], 0, len(a[i])-1, &result)
	}
	return result
}

func generateAllAnagrams(s string, l, r int, result *[]string) {
	if l == r {
		newS := strings.Clone(s)
		*result = append(*result, newS)
	} else {
		for i := l; i <= r; i++ {
			swapChar(&s, l, i)
			generateAllAnagrams(s, l+1, r, result)
			swapChar(&s, l, i)
		}
	}
}
func swapChar(s *string, i, j int) {
	n := len(*s)
	arr := []byte(*s)
	if i >= 0 && j < n {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
	*s = string(arr)
}
