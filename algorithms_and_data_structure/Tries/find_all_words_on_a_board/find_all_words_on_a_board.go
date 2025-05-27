package main

import (
	"fmt"
)

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func findAllWordsOnBoard(board [][]byte, words []string) []string {
	root := NewTrieNode()

	// Построение Trie
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if _, ok := node.children[c]; !ok {
				node.children[c] = NewTrieNode()
			}
			node = node.children[c]
		}
		node.word = word
	}

	res := []string{}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if node, ok := root.children[board[r][c]]; ok {
				dfs(board, r, c, node, &res)
			}
		}
	}
	return res
}

func dfs(board [][]byte, r, c int, node *TrieNode, res *[]string) {
	if node.word != "" {
		*res = append(*res, node.word)
		node.word = "" // избегаем дубликатов
	}

	temp := board[r][c]
	board[r][c] = '#'

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, d := range directions {
		newR, newC := r+d[0], c+d[1]
		if isWithinBounds(newR, newC, board) {
			nextChar := board[newR][newC]
			if nextNode, ok := node.children[nextChar]; ok {
				dfs(board, newR, newC, nextNode, res)
			}
		}
	}

	board[r][c] = temp
}

func isWithinBounds(r, c int, board [][]byte) bool {
	return r >= 0 && r < len(board) && c >= 0 && c < len(board[0])
}

// Пример использования
func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}

	found := findAllWordsOnBoard(board, words)
	fmt.Println(found) // пример: [oath eat]
}
