package main

import "fmt"

/*
Q1 ) Given a log of website requests, where each request entry contains (time, customerId, page visited),
find the top 3-page sequence visited.

For example, if we have two customers, and we log
CustomerA visiting page A->B->C->D->E and
CustomerB visiting E->B->C->D->A,

output: B->C->D
*/
func main() {
	visitedSequences := []string{"ABCDE", "EBCDA", "DEAB"}
	fmt.Println(findMostVisitedSequence(visitedSequences))
}

func findMostVisitedSequence(sequences []string) string {
	var res string
	max := -1
	dp := make(map[string]int)
	for i := 0; i < len(sequences); i++ {
		seq := sequences[i]
		for j := 0; j < len(seq)-3; j++ {
			v, _ := dp[seq[j:j+3]]
			dp[seq[j:j+3]] = v + 1
			if max < v+1 {
				max = v + 1
				res = seq[j : j+3]
			}

		}
	}
	return res
}
