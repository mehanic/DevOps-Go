package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func kthSmallestNumberInBSTIterative(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    node := root
    for len(stack) > 0 || node != nil {
        // Move to the leftmost node, pushing nodes on stack
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        // Pop node from stack
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k--
        if k == 0 {
            return node.Val
        }
        // Move to right subtree
        node = node.Right
    }
    return -1 // If k is invalid
}
