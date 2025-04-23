package main

import (
	"fmt"
)
//1. Depth First Search
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
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
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


//2. Iterative DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree1(p *TreeNode, q *TreeNode) bool {
    type Pair struct {
        first, second *TreeNode
    }
    
    stack := []Pair{{p, q}}
    
    for len(stack) > 0 {
        lastIdx := len(stack) - 1
        node1, node2 := stack[lastIdx].first, stack[lastIdx].second
        stack = stack[:lastIdx]
        
        if node1 == nil && node2 == nil {
            continue
        }
        if node1 == nil || node2 == nil || node1.Val != node2.Val {
            return false
        }
        
        stack = append(stack, Pair{node1.Right, node2.Right})
        stack = append(stack, Pair{node1.Left, node2.Left})
    }
    
    return true
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
 func isSameTree2(p *TreeNode, q *TreeNode) bool {
    queue1 := []*TreeNode{p}
    queue2 := []*TreeNode{q}

    for len(queue1) > 0 && len(queue2) > 0 {
        for i := len(queue1); i > 0; i-- {
            nodeP := queue1[0]
            nodeQ := queue2[0]
            queue1 = queue1[1:]
            queue2 = queue2[1:]

            if nodeP == nil && nodeQ == nil {
                continue
            }
            if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
                return false
            }

            queue1 = append(queue1, nodeP.Left, nodeP.Right)
            queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
        }
    }

    return len(queue1) == 0 && len(queue2) == 0
}

