package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	n := len(emails)
	dp := make(map[string]bool)

	for i := 0; i < n; i++ {
		email := emails[i]
		localDomail := strings.Split(email, "@")
		localName := strings.ReplaceAll(localDomail[0], ".", "")
		localName = strings.Split(localName, "+")[0]
		newEmail := fmt.Sprintf("%s%s%s", localName+"@"+localDomail[1])
		if !dp[newEmail] {
			dp[newEmail] = true
		}
	}
	return len(dp)
}
