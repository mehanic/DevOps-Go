/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftHeight := maxHeight(root.Left)
    rightHeight := maxHeight(root.Right)
    diameter := leftHeight + rightHeight
    
    sub := max(diameterOfBinaryTree(root.Left), 
              diameterOfBinaryTree(root.Right))
    
    return max(diameter, sub)
}

func maxHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return 1 + max(maxHeight(root.Left), maxHeight(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}