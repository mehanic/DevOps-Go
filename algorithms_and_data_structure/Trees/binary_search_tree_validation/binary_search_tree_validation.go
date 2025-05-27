package main

import (
	"math"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binarySearchTreeValidation(root *TreeNode) bool {
	return isWithinBounds(root, math.Inf(-1), math.Inf(1))
}

func isWithinBounds(node *TreeNode, lowerBound, upperBound float64) bool {
	if node == nil {
		return true
	}
	val := float64(node.Val)
	if val <= lowerBound || val >= upperBound {
		return false
	}
	if !isWithinBounds(node.Left, lowerBound, val) {
		return false
	}
	return isWithinBounds(node.Right, val, upperBound)
}
