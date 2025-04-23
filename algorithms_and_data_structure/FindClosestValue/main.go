package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, math.MaxInt32)
}

func (tree *BST) findClosestValue(target, closest int) int {
	// Check if current value is closer than the previous closest value
	if absdiff(target, closest) > absdiff(target, tree.Value) {
		closest = tree.Value
	}

	// If target is less than the current node value, go left
	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosestValue(target, closest)
		// If target is greater than the current node value, go right
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosestValue(target, closest)
	}
	// Return the closest value found so far
	return closest
}

// Helper function to calculate absolute difference
func absdiff(a, b int) int {
	out := math.Abs(float64(a) - float64(b))
	return int(out)
}

func main() {
	// Example: Constructing a sample Binary Search Tree
	root := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
			},
			Right: &BST{
				Value: 5,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
			},
			Right: &BST{
				Value: 22,
			},
		},
	}

	// Testing with different targets
	targets := []int{12, 3, 18, 10}

	for _, target := range targets {
		fmt.Printf("Closest value to %d: %d\n", target, root.FindClosestValue(target))
	}
}
