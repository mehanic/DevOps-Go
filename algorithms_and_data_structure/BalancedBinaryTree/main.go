package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to check balance and compute height
func checkBalance(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	// Check left subtree
	leftBalanced, leftHeight := checkBalance(root.Left)
	// Check right subtree
	rightBalanced, rightHeight := checkBalance(root.Right)

	// Current node is balanced if left and right subtrees are balanced
	// and their height difference is no more than 1
	balanced := leftBalanced && rightBalanced && math.Abs(float64(leftHeight-rightHeight)) <= 1
	height := max(leftHeight, rightHeight) + 1

	return balanced, height
}

// Function to check if the tree is height-balanced
func isBalanced(root *TreeNode) bool {
	balanced, _ := checkBalance(root)
	return balanced
}

// Helper function to get the max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example Usage
func main() {
	// Creating the tree: [1,2,3,nil,nil,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}

	// Checking if the tree is balanced
	fmt.Println("Is the tree balanced?", isBalanced(root)) // Output: true

	// Creating an unbalanced tree: [1,2,3,nil,nil,4,nil,5]
	root.Right.Left.Left = &TreeNode{Val: 5}

	// Checking again
	fmt.Println("Is the tree balanced?", isBalanced(root)) // Output: false
}
