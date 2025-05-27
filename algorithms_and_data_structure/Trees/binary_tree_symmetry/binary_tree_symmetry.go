package main

// TreeNode definition
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func binaryTreeSymmetry(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return compareTrees(root.Left, root.Right)
}

func compareTrees(node1, node2 *TreeNode) bool {
    // Both nil means symmetric
    if node1 == nil && node2 == nil {
        return true
    }
    // One nil means not symmetric
    if node1 == nil || node2 == nil {
        return false
    }
    // Values mismatch means not symmetric
    if node1.Val != node2.Val {
        return false
    }
    // Recursively compare left subtree of node1 with right subtree of node2
    if !compareTrees(node1.Left, node2.Right) {
        return false
    }
    // Recursively compare right subtree of node1 with left subtree of node2
    return compareTrees(node1.Right, node2.Left)
}
