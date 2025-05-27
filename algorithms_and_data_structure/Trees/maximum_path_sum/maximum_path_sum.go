package main

import "math"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
    maxSum = math.MinInt32
    maxPathSumHelper(root)
    return maxSum
}

func maxPathSumHelper(node *TreeNode) int {
    if node == nil {
        return 0
    }
    leftSum := max(maxPathSumHelper(node.Left), 0)
    rightSum := max(maxPathSumHelper(node.Right), 0)

    // Update global maxSum if current path sum is greater
    currentSum := node.Val + leftSum + rightSum
    if currentSum > maxSum {
        maxSum = currentSum
    }

    // Return max sum of one side plus current node
    return node.Val + max(leftSum, rightSum)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
