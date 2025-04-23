
// Invert Binary Tree
// You are given the root of a binary tree root. Invert the binary tree and return its root.

// Example 1:



// Input: root = [1,2,3,4,5,6,7]

// Output: [1,3,2,7,6,5,4]
// Example 2:



// Input: root = [3,2,1]

// Output: [3,1,2]
// Example 3:

// Input: root = []

// Output: []
// Constraints:

// 0 <= The number of nodes in the tree <= 100.
// -100 <= Node.val <= 100



package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to invert a binary tree (Recursive)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right subtrees
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert left and right subtrees
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Helper function to print tree in level order (Breadth-First Traversal)
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			fmt.Print(node.Val, " ")
			queue = append(queue, node.Left, node.Right)
		} else {
			fmt.Print("null ")
		}
	}
	fmt.Println()
}

// Helper function to create a binary tree from an array representation
func createTreeFromArray(arr []int, index int) *TreeNode {
	if index >= len(arr) || arr[index] == -1 {
		return nil
	}

	node := &TreeNode{Val: arr[index]}
	node.Left = createTreeFromArray(arr, 2*index+1)
	node.Right = createTreeFromArray(arr, 2*index+2)

	return node
}

// Test Cases
func main() {
	// Example 1: Input tree [1,2,3,4,5,6,7]
	root1 := createTreeFromArray([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	fmt.Print("Original Tree: ")
	printTree(root1)
	root1 = invertTree(root1)
	fmt.Print("Inverted Tree: ")
	printTree(root1)

	// Example 2: Input tree [3,2,1]
	root2 := createTreeFromArray([]int{3, 2, 1}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root2)
	root2 = invertTree(root2)
	fmt.Print("Inverted Tree: ")
	printTree(root2)

	// Example 3: Empty tree
	root3 := createTreeFromArray([]int{}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root3)
	root3 = invertTree(root3)
	fmt.Print("Inverted Tree: ")
	printTree(root3)
}
