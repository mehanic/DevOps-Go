package main

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type Level struct {
	Root  *BinaryTree
	Depth int
}

// Function to calculate the sum of node depths
func NodeDepths(root *BinaryTree) int {
	sumOfDepths := 0
	stack := []Level{{Root: root, Depth: 0}} // Initialize stack with root node and depth 0
	var top Level

	// Iterative depth-first traversal using stack
	for len(stack) > 0 {
		// Pop the top node from the stack
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		node, depth := top.Root, top.Depth

		// If the node is nil, skip it
		if node == nil {
			continue
		}

		// Add the current depth to the sum of depths
		sumOfDepths += depth

		// Push left and right children (if any) to the stack with incremented depth
		stack = append(stack, Level{Root: node.Left, Depth: depth + 1})
		stack = append(stack, Level{Root: node.Right, Depth: depth + 1})
	}

	// Return the sum of depths
	return sumOfDepths
}

func main() {
	// Example Binary Tree
	root := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
		},
	}

	// Calculate the sum of node depths
	result := NodeDepths(root)
	fmt.Println("Sum of Node Depths:", result)
}
