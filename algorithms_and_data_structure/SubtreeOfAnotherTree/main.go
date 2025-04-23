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

// Function to check if a tree is a subtree
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Function to check if two trees are identical
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

// Helper function to print the result
func printResult(root, subRoot *TreeNode) {
	fmt.Println(isSubtree(root, subRoot))
}

// Example usage
func main() {
	// Example 1: subRoot is a valid subtree of root
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	subRoot1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	printResult(root1, subRoot1) // Output: true

	// Example 2: subRoot is not a valid subtree of root
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	subRoot2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	printResult(root2, subRoot2) // Output: false

	// Example 3: subRoot is the same as root (edge case)
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	subRoot3 := root3
	printResult(root3, subRoot3) // Output: true

	// Example 4: subRoot does not exist in root
	root4 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}}
	subRoot4 := &TreeNode{Val: 5}
	printResult(root4, subRoot4) // Output: false

	// Example 5: subRoot is empty
	var root5 *TreeNode
	var subRoot5 *TreeNode
	printResult(root5, subRoot5) // Output: true (an empty tree is a subtree of an empty tree)
}
