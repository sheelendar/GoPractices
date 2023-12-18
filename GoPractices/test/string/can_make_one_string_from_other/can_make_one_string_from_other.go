package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	m := len(ransomNote)
	n := len(magazine)
	mapping := make(map[byte]int)
	for i := 0; i < m; i++ {
		mapping[ransomNote[i]]++
	}
	for i := 0; i < n; i++ {
		mapping[magazine[i]]--
	}
	for _, count := range mapping {
		if count > 0 {
			return false
		}
	}
	return true
}
