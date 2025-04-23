package main

import (
	"container/list"
	"fmt"
	//"github.com/emirpasic/gods/queues/linkedlistqueue" // Importing a queue implementation
)

//1. Recursive DFS
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



//2. Iterative DFS (Stack)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    stack := list.New()
    stack.PushBack([]interface{}{root, 1})
    res := 0
    
    for stack.Len() > 0 {
        back := stack.Back()
        stack.Remove(back)
        pair := back.Value.([]interface{})
        node := pair[0].(*TreeNode)
        depth := pair[1].(int)
        
        if node != nil {
            res = max(res, depth)
            stack.PushBack([]interface{}{node.Left, depth + 1})
            stack.PushBack([]interface{}{node.Right, depth + 1})
        }
    }
    
    return res
}

func max1(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//3. Breadth First Search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  func maxDepth3(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
	
// 	q := linkedlistqueue.New()
// 	q.Enqueue(root)
// 	level := 0
	
// 	for !q.Empty() {
// 		size := q.Size()
		
// 		for i := 0; i < size; i++ {
// 			val, _ := q.Dequeue()
// 			node := val.(*TreeNode)
			
// 			if node.Left != nil {
// 				q.Enqueue(node.Left)
// 			}
// 			if node.Right != nil {
// 				q.Enqueue(node.Right)
// 			}
// 		}
// 		level++
// 	}
	
// 	return level
//  }


//-------------

// // Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// // Function to compute max depth using BFS (level order traversal)
// func maxDepth3(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	q := linkedlistqueue.New()
// 	q.Enqueue(root)
// 	level := 0

// 	for !q.Empty() {
// 		size := q.Size()

// 		for i := 0; i < size; i++ {
// 			val, _ := q.Dequeue()
// 			node := val.(*TreeNode)

// 			if node.Left != nil {
// 				q.Enqueue(node.Left)
// 			}
// 			if node.Right != nil {
// 				q.Enqueue(node.Right)
// 			}
// 		}
// 		level++
// 	}

// 	return level
// }

// // Helper function to create a binary tree from an array
// func createTreeFromArray(arr []int, index int) *TreeNode {
// 	if index >= len(arr) || arr[index] == -1 {
// 		return nil
// 	}
// 	node := &TreeNode{Val: arr[index]}
// 	node.Left = createTreeFromArray(arr, 2*index+1)
// 	node.Right = createTreeFromArray(arr, 2*index+2)
// 	return node
// }

// // Main function with test cases
// func main() {
// 	// Example 1: Balanced tree
// 	root1 := createTreeFromArray([]int{1, 2, 3, -1, -1, 4, 5}, 0)
// 	fmt.Println("Maximum Depth of Example 1:", maxDepth3(root1)) // Output: 3

// 	// Example 2: Empty tree
// 	root2 := createTreeFromArray([]int{}, 0)
// 	fmt.Println("Maximum Depth of Example 2:", maxDepth3(root2)) // Output: 0

// 	// Example 3: Skewed tree
// 	root3 := createTreeFromArray([]int{10, 20, -1, 30}, 0)
// 	fmt.Println("Maximum Depth of Example 3:", maxDepth3(root3)) // Output: 3
// }
