package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // return True
	fmt.Println(trie.Search("app"))     // return False
	fmt.Println(trie.StartsWith("app")) // return True
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // return True
}

type Trie struct {
	adj   map[byte]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{adj: make(map[byte]*Trie)}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	curr := this
	for i := 0; i < n; i++ {
		if _, ok := curr.adj[word[i]]; !ok {
			trie := Constructor()
			curr.adj[word[i]] = &trie
		}
		curr = curr.adj[word[i]]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	curr := this
	for i := 0; i < n; i++ {
		if _, ok := curr.adj[word[i]]; !ok {
			return false
		}
		curr = curr.adj[word[i]]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	curr := this
	for i := 0; i < n; i++ {
		if _, ok := curr.adj[prefix[i]]; !ok {
			return false
		}
		curr = curr.adj[prefix[i]]
	}
	return true
}
