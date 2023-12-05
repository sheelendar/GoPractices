package main

import (
	"fmt"
	"sort"
)

/*
Search Suggestion System
For an array of n strings products and a word to search, search, design a system that, when each character of the searched word is typed, suggests at most three product names from the products array. The suggested products should share a common prefix with the searched word. If more than three products exist with a common prefix, report the three product names that appear first in lexicographical order.
Return the suggested products, which will be a list of lists after each character of the searched word is typed.
Note: A string x is considered lexicographically smaller than another string y if × will occur before y in the English dictionary.
Example
Suppose n = 5, products = ["carpet", "cart", "car"
"camera", "crate"], and search = "camera".
Search
Prefix					Matches      									Lexicographically Smallest 3
c			["carpet", "cart","car", "camera","crate"]				["camera", "car","carpet"]
ca			['carpet", "cart","car", "camera"]						['camera", "car","carpet"]
cam			['camera"]												['camera"]
came		['camera"]												['camera"]
camer		['camera"]												['camera"]
camera		['camera"] 												['camera"]

Hence the answer is [['camera" "car" "carpet"], ['camera"
"car", "carpet"], ["camera"], ["camera"], ["camera"], ['camera"]].

Function Description
Complete the function getProductSuggestions in the editor below.
getProductSuggestions has the following parameter(s):
string products[n]: the list of products string search: a string

Returns
string[n]I: for each prefix of the string, return a maximum of three lexicographically smallest words with a common prefix.
Constraints
• 1 <n <1000
• 1 ≤ length of products[i] ≤ 500
• 1 ≤ sum(length of productsil) ≤ 5 × 105
• All the strings of products are unique.
• products[i] consists of lowercase English letters only.
• 1 ≤ length of search ≤1000
• The searched text consists of lowercase English letters only.
*/
func main() {
	products := []string{"carpet", "cart", "car", "camera", "crate"}
	search := "camera"
	fmt.Println(getProductSuggestions(products, search))
}

func getProductSuggestions(products []string, search string) [][]string {
	dp := make(map[string][]string)
	n := len(products)
	var p []string
	remo := make(map[string]bool)
	for i := 0; i < n; i++ {
		if !remo[products[i]] {
			p = append(p, products[i])
			remo[products[i]] = true
		}
	}
	for i := 0; i < n; i++ {
		size := len(p[i])
		str := p[i]
		for j := 1; j <= size; j++ {
			dp[str[:j]] = append(dp[str[:j]], str)
		}
	}
	var result [][]string
	for i := 1; i <= len(search); i++ {
		str := search[:i]
		list := dp[str]
		if len(list) == 0 {
			continue
		}
		sort.Strings(list)
		if len(list) >= 3 {
			result = append(result, []string{list[0], list[1], list[2]})
		} else {
			result = append(result, list)
		}
	}
	return result

}
