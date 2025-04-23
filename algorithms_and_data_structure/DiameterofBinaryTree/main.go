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

// Helper function to calculate height and update max diameter
func heightAndDiameter(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	leftHeight := heightAndDiameter(root.Left, maxDiameter)
	rightHeight := heightAndDiameter(root.Right, maxDiameter)

	// Update the maximum diameter found so far
	*maxDiameter = max(*maxDiameter, leftHeight+rightHeight)

	// Return height of the current node
	return 1 + max(leftHeight, rightHeight)
}

// Function to find the diameter of a binary tree
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	heightAndDiameter(root, &maxDiameter)
	return maxDiameter
}

// Utility function to create a binary tree from an array
func createTreeFromArray(arr []int, index int) *TreeNode {
	if index >= len(arr) || arr[index] == -1 {
		return nil
	}
	node := &TreeNode{Val: arr[index]}
	node.Left = createTreeFromArray(arr, 2*index+1)
	node.Right = createTreeFromArray(arr, 2*index+2)
	return node
}

// Helper function to get max of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function with test cases
func main() {
	// Example 1: Tree with a long path
	root1 := createTreeFromArray([]int{1, -1, 2, 3, 4, 5}, 0)
	fmt.Println("Diameter of Example 1:", diameterOfBinaryTree(root1)) // Output: 3

	// Example 2: Balanced Tree
	root2 := createTreeFromArray([]int{1, 2, 3}, 0)
	fmt.Println("Diameter of Example 2:", diameterOfBinaryTree(root2)) // Output: 2

	// Example 3: Single node tree
	root3 := createTreeFromArray([]int{1}, 0)
	fmt.Println("Diameter of Example 3:", diameterOfBinaryTree(root3)) // Output: 0
}
