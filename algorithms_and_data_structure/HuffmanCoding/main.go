package main

import (
	"fmt"
	"sort"
	"strings"
)

// NodeTree defines a node in the Huffman tree
type NodeTree struct {
	Left   interface{}
	Right  interface{}
	Weight int
}

// Pair is used to hold a character/node with its frequency
type Pair struct {
	CharOrNode interface{}
	Freq       int
}

// Recursively generate Huffman codes
func huffmanCodeTree(node interface{}, binString string, codeMap map[string]string) {
	switch n := node.(type) {
	case string:
		codeMap[n] = binString
	case *NodeTree:
		huffmanCodeTree(n.Left, binString+"0", codeMap)
		huffmanCodeTree(n.Right, binString+"1", codeMap)
	}
}

// Reverse map: value to key
func reverseCodeMap(codeMap map[string]string) map[string]string {
	reversed := make(map[string]string)
	for k, v := range codeMap {
		reversed[v] = k
	}
	return reversed
}

// Encode input string to binary using Huffman codes
func encode(input string, codeMap map[string]string) string {
	var encoded strings.Builder
	for _, c := range input {
		encoded.WriteString(codeMap[string(c)])
	}
	return encoded.String()
}

// Decode binary string using Huffman tree
func decode(encoded string, root interface{}) string {
	var decoded strings.Builder
	node := root

	for _, bit := range encoded {
		if tree, ok := node.(*NodeTree); ok {
			if bit == '0' {
				node = tree.Left
			} else {
				node = tree.Right
			}
		}

		if char, ok := node.(string); ok {
			decoded.WriteString(char)
			node = root
		}
	}
	return decoded.String()
}

func main() {
	input := "abhsbaa"

	// Frequency calculation
	freq := make(map[string]int)
	for _, c := range input {
		char := string(c)
		freq[char]++
	}

	// Convert to slice of Pair
	var nodes []Pair
	for char, f := range freq {
		nodes = append(nodes, Pair{CharOrNode: char, Freq: f})
	}

	// Sort descending by frequency
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Freq > nodes[j].Freq
	})

	// Build Huffman Tree
	for len(nodes) > 1 {
		// Take two lowest freq elements
		last := len(nodes)
		left := nodes[last-1]
		right := nodes[last-2]

		newNode := &NodeTree{
			Left:   left.CharOrNode,
			Right:  right.CharOrNode,
			Weight: left.Freq + right.Freq,
		}

		nodes = nodes[:last-2]
		nodes = append(nodes, Pair{CharOrNode: newNode, Freq: newNode.Weight})

		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Freq > nodes[j].Freq
		})
	}

	// Generate Huffman codes
	codeMap := make(map[string]string)
	huffmanCodeTree(nodes[0].CharOrNode, "", codeMap)

	// Display the Huffman codes
	fmt.Println(" Char | Huffman Code")
	fmt.Println("----------------------")
	for char, code := range codeMap {
		fmt.Printf("  %-3s | %s\n", char, code)
	}

	// Encode
	encoded := encode(input, codeMap)
	fmt.Println("\nOriginal String:", input)
	fmt.Println("Encoded Binary :", encoded)

	// Decode
	decoded := decode(encoded, nodes[0].CharOrNode)
	fmt.Println("Decoded String :", decoded)
}
