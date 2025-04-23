//1. Breadth First Search

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

// Iterative BFS Invert Tree Function
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap left and right subtrees
		node.Left, node.Right = node.Right, node.Left

		// Enqueue left and right children
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// Helper function to print the tree level order
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




//2. Depth First Search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func invertTree2(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)

    return root
}


func main1() {
	// Example 1: Input tree [1,2,3,4,5,6,7]
	root1 := createTreeFromArray([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	fmt.Print("Original Tree: ")
	printTree(root1)
	root1 = invertTree2(root1)
	fmt.Print("Inverted Tree: ")
	printTree(root1)

	// Example 2: Input tree [3,2,1]
	root2 := createTreeFromArray([]int{3, 2, 1}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root2)
	root2 = invertTree2(root2)
	fmt.Print("Inverted Tree: ")
	printTree(root2)

	// Example 3: Empty tree
	root3 := createTreeFromArray([]int{}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root3)
	root3 = invertTree2(root3)
	fmt.Print("Inverted Tree: ")
	printTree(root3)
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
 func invertTree3(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left, root.Right = root.Right, root.Left

    invertTree(root.Left)
    invertTree(root.Right)

    return root
}



func main2() {
	// Example 1: Input tree [1,2,3,4,5,6,7]
	root1 := createTreeFromArray([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	fmt.Print("Original Tree: ")
	printTree(root1)
	root1 = invertTree3(root1)
	fmt.Print("Inverted Tree: ")
	printTree(root1)

	// Example 2: Input tree [3,2,1]
	root2 := createTreeFromArray([]int{3, 2, 1}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root2)
	root2 = invertTree3(root2)
	fmt.Print("Inverted Tree: ")
	printTree(root2)

	// Example 3: Empty tree
	root3 := createTreeFromArray([]int{}, 0)
	fmt.Print("\nOriginal Tree: ")
	printTree(root3)
	root3 = invertTree3(root3)
	fmt.Print("Inverted Tree: ")
	printTree(root3)
}
