package main

import "fmt"

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

	Trie() Initializes the trie object.
	void insert(String word) Inserts the string word into the trie.
	boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
	boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
*/
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
