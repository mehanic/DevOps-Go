package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type WordDictionary struct {
	root *TrieNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: NewTrieNode()}
}

func (wd *WordDictionary) Insert(word string) {
	node := wd.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrieNode()
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.searchHelper(word, 0, wd.root)
}

func (wd *WordDictionary) searchHelper(word string, index int, node *TrieNode) bool {
	for i := index; i < len(word); i++ {
		c := rune(word[i])
		if c == '.' {
			// wildcard: ищем рекурсивно по всем детям
			for _, child := range node.children {
				if wd.searchHelper(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			if nextNode, ok := node.children[c]; ok {
				node = nextNode
			} else {
				return false
			}
		}
	}
	return node.isWord
}

// Пример использования
func main() {
	wd := NewWordDictionary()
	wd.Insert("bad")
	wd.Insert("dad")
	wd.Insert("mad")

	fmt.Println(wd.Search("pad"))  // false
	fmt.Println(wd.Search("bad"))  // true
	fmt.Println(wd.Search(".ad"))  // true
	fmt.Println(wd.Search("b.."))  // true
}
