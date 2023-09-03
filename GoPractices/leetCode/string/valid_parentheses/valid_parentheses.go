package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isPalindrome("0P"))
}

func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	if n == 1 {
		return false
	}
	data := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	var stack []byte
	for i := 0; i < n; i++ {
		if v, ok := data[s[i]]; ok {
			stack = append(stack, v)
		} else {
			l := len(stack)
			if l > 0 && stack[l-1] == s[i] {
				stack = stack[:l-1]

			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	n := len(s)
	if n == 0 || n == 1 {
		return true
	}
	arr := []string{",", ";", ":", "?", "!", "'", ".", " "}

	for i := 0; i < len(arr); i++ {
		s = strings.ReplaceAll(s, arr[i], "")
	}

	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] == s[j] || ((s[i] > 64) && (s[i] < 97) && s[i]+32 == s[j]) || ((s[j]) > 64 && (s[j] < 97) && s[j]+32 == s[i]) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
