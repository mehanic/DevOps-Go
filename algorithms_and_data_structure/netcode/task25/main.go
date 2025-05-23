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
//1. Brute Force
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


//2. Depth First Search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree1(root *TreeNode) int {
    res := 0
    
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        
        left := dfs(root.Left)
        right := dfs(root.Right)
        res = max(res, left + right)
        
        return 1 + max(left, right)
    }
    
    dfs(root)
    return res
}

func max1(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//3. Iterative DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	stack := linkedliststack.New()
	stack.Push(root)
	mp := make(map[*TreeNode][]int)
	mp[nil] = []int{0, 0}
	
	for !stack.Empty() {
		val, _ := stack.Peek()
		node := val.(*TreeNode)
		
		if node.Left != nil && len(mp[node.Left]) == 0 {
			stack.Push(node.Left)
		} else if node.Right != nil && len(mp[node.Right]) == 0 {
			stack.Push(node.Right)
		} else {
			stack.Pop()
			
			leftHeight := mp[node.Left][0]
			leftDiameter := mp[node.Left][1]
			rightHeight := mp[node.Right][0]
			rightDiameter := mp[node.Right][1]
			
			height := 1 + max(leftHeight, rightHeight)
			diameter := max(leftHeight+rightHeight, 
						  max(leftDiameter, rightDiameter))
			
			mp[node] = []int{height, diameter}
		}
	}
	
	return mp[root][1]
 }
 
 func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
 }