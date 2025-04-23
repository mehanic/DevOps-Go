package main

import "fmt"

//1. Brute Force
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    left := height(root.Left)
    right := height(root.Right)
    if abs(left-right) > 1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}

    fmt.Println(isBalanced(root)) // Output: true
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
 func isBalanced1(root *TreeNode) bool {
    return dfs(root).balanced
}

type Result struct {
    balanced bool
    height   int
}

func dfs(root *TreeNode) Result {
    if root == nil {
        return Result{true, 0}
    }
    
    left := dfs(root.Left)
    right := dfs(root.Right)
    
    balanced := left.balanced && right.balanced && abs(left.height - right.height) <= 1
    return Result{balanced, 1 + max(left.height, right.height)}
}

func max1(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs1(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


///3. Iterative DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced2(root *TreeNode) bool {
    stack := []*TreeNode{}
    node := root
    last := (*TreeNode)(nil)
    depths := make(map[*TreeNode]int)

    for len(stack) > 0 || node != nil {
        if node != nil {
            stack = append(stack, node)
            node = node.Left
        } else {
            node = stack[len(stack)-1]
            if node.Right == nil || last == node.Right {
                stack = stack[:len(stack)-1]
                left := depths[node.Left]
                right := depths[node.Right]

                if abs(left-right) > 1 {
                    return false
                }

                depths[node] = 1 + max(left, right)
                last = node
                node = nil
            } else {
                node = node.Right
            }
        }
    }
    return true
}

func abs2(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max2(a, b int) int {
    if a > b {
        return a
    }
    return b
}