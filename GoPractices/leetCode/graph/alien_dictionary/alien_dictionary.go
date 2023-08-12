package main

import "fmt"

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
	for k, _ := range rela {
		DFS(k, rela, &stack, visited)
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
