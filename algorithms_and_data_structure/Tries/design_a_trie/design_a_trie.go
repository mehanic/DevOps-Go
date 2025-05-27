package main

import "fmt"

// TrieNode представляет узел дерева
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// Trie представляет префиксное дерево
type Trie struct {
	root *TrieNode
}

// NewTrie возвращает новый Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

// Insert добавляет слово в Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		if _, exists := node.children[c]; !exists {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isWord = true
}

// Search проверяет, существует ли слово в Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		if _, exists := node.children[c]; !exists {
			return false
		}
		node = node.children[c]
	}
	return node.isWord
}

// HasPrefix проверяет, существует ли префикс в Trie
func (t *Trie) HasPrefix(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		if _, exists := node.children[c]; !exists {
			return false
		}
		node = node.children[c]
	}
	return true
}

// Пример использования
func main() {
	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println(trie.Search("hello"))   // true
	fmt.Println(trie.Search("hell"))    // false
	fmt.Println(trie.HasPrefix("hell")) // true
	fmt.Println(trie.HasPrefix("hi"))   // false
}
