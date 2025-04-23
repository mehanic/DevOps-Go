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

// Function to check if two binary trees are identical
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Helper function to print the result
func printResult(p, q *TreeNode) {
	fmt.Println(isSameTree(p, q))
}

// Example usage
func main() {
	// Example 1: Identical trees
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	printResult(p1, q1) // Output: true

	// Example 2: Different structures
	p2 := &TreeNode{Val: 4, Left: &TreeNode{Val: 7}}
	q2 := &TreeNode{Val: 4, Right: &TreeNode{Val: 7}}
	printResult(p2, q2) // Output: false

	// Example 3: Same structure but different values
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}
	printResult(p3, q3) // Output: false

	// Example 4: One tree is empty
	var p4 *TreeNode
	q4 := &TreeNode{Val: 1}
	printResult(p4, q4) // Output: false

	// Example 5: Both trees are empty
	var p5, q5 *TreeNode
	printResult(p5, q5) // Output: true
}
