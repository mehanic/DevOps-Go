package main

import (
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balancedBinaryTreeValidation(root *TreeNode) bool {
	return getHeightImbalance(root) != -1
}

func getHeightImbalance(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeightImbalance(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeightImbalance(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
