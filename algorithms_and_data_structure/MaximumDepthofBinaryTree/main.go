// Maximum Depth of Binary Tree
// Given the root of a binary tree, return its depth.

// The depth of a binary tree is defined as the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:



// Input: root = [1,2,3,null,null,4]

// Output: 3
// Example 2:

// Input: root = []

// Output: 0
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

// Recursive DFS function to find the maximum depth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

// Helper function to get the max value between two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to create a binary tree from an array
func createTreeFromArray(arr []int, index int) *TreeNode {
	if index >= len(arr) || arr[index] == -1 {
		return nil
	}
	node := &TreeNode{Val: arr[index]}
	node.Left = createTreeFromArray(arr, 2*index+1)
	node.Right = createTreeFromArray(arr, 2*index+2)
	return node
}

// Main function with test cases
func main() {
	// Example 1: Input tree [1,2,3,null,null,4]
	root1 := createTreeFromArray([]int{1, 2, 3, -1, -1, 4}, 0)
	fmt.Println("Maximum Depth of Example 1:", maxDepth(root1)) // Output: 3

	// Example 2: Empty tree []
	root2 := createTreeFromArray([]int{}, 0)
	fmt.Println("Maximum Depth of Example 2:", maxDepth(root2)) // Output: 0

	// Example 3: A deeper tree [1,2,3,4,5,null,6]
	root3 := createTreeFromArray([]int{1, 2, 3, 4, 5, -1, 6}, 0)
	fmt.Println("Maximum Depth of Example 3:", maxDepth(root3)) // Output: 3
}
