package main

import "fmt"

/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

	WordDictionary() Initializes the object.
	void addWord(word) Adds word to the data structure, it can be matched later.
	bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
*/

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad")) // return False
	fmt.Println(wordDictionary.Search("bad")) // return True
	fmt.Println(wordDictionary.Search(".ad")) // return True
	fmt.Println(wordDictionary.Search("b..")) // return True
}

type WordDictionary struct {
	adj   map[byte]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{adj: make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	n := len(word)
	for i := 0; i < n; i++ {
		if _, ok := curr.adj[word[i]]; !ok {
			node := Constructor()
			curr.adj[word[i]] = &node
		}
		curr = curr.adj[word[i]]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this
	n := len(word)

	for i := 0; i < n; i++ {
		if word[i] == '.' {
			res := false
			for _, a := range curr.adj {
				res = a.Search(word[i+1:])
				if res {
					break
				}
			}
			return res
		} else {
			if _, ok := curr.adj[word[i]]; ok {
				curr = curr.adj[word[i]]
			} else {
				return false
			}
		}
	}
	return curr.isEnd
}
