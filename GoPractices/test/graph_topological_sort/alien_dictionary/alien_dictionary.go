package main

import "fmt"

/*
Given a sorted dictionary of an alien language having N words and k starting alphabets of standard dictionary.
Find the order of characters in the alien language.

Note: Many orders may be possible for a particular test case, thus you may return any valid order and output will be 1 if
the order of string returned by the function is correct else 0 denoting incorrect string
Example 1:
Input:
N = 5, K = 4
dict = {"baa","abcd","abca","cab","cad"}
Output:
1
Explanation:
Here order of characters is
'b', 'd', 'a', 'c' Note that words are sorted
and in the given language "baa" comes before
"abcd", therefore 'b' is before 'a' in output.
Similarly we can find other order
*/
func main() {
	N := 5
	K := 4
	dict := []string{"baa", "abcd", "abca", "cab", "cad"}
	findOrderOfLetters(dict, N, K)
}

func findOrderOfLetters(dict []string, N, K int) {
	rela := make(map[byte][]byte)

	// making a dependency graph here and take it into rela
	for i := 1; i < N; i++ {
		a, b := checkRelations(dict[i-1], dict[i])
		if _, ok := rela[a]; ok {
			rela[a] = append(rela[a], b)
		} else {
			rela[a] = []byte{b}
		}
	}
	visited := make(map[byte]bool)
	var stack []byte
	// call dfs as topological sort and collect output into stack to print correct order.
	for node, _ := range rela {
		DFS(node, rela, &stack, visited)
	}
	//displayRelation(rela)
	fmt.Println()
	sep := ""
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Print(sep)
		fmt.Printf("%c", stack[i])
		sep = " -> "
	}
}

// compare both word and find out a relation b/w word.
// str1 always small whenever both strings are not equal then ste1[i] is less than str2[j]
func checkRelations(str1, str2 string) (byte, byte) {
	n1 := len(str1)
	n2 := len(str2)
	i := 0
	j := 0
	for i < n1 && j < n2 {
		if str1[i] != str2[j] {
			break
		}
		i++
		j++
	}
	return str1[i], str2[j]
}

func DFS(n byte, rela map[byte][]byte, stack *[]byte, visited map[byte]bool) {

	if visited[n] {
		return
	}
	adj := rela[n]
	for j := 0; j < len(adj); j++ {
		if !visited[adj[j]] {
			DFS(adj[j], rela, stack, visited)
		}
	}
	visited[n] = true
	*stack = append(*stack, n)
}

// this is to check if graph is formed correctly or not
func displayRelation(rela map[byte][]byte) {
	for k, v := range rela {
		fmt.Printf("%c", k)
		fmt.Print(" -> ")
		for i := 0; i < len(v); i++ {
			fmt.Printf("%c", v[i])
			fmt.Print(", ")
		}
		fmt.Println()
	}
}
