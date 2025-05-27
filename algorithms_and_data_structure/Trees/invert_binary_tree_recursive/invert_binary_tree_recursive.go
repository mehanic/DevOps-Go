package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func invertBinaryTreeRecursive(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // Swap left and right children
    root.Left, root.Right = root.Right, root.Left
    // Recursively invert left and right subtrees
    invertBinaryTreeRecursive(root.Left)
    invertBinaryTreeRecursive(root.Right)
    return root
}
