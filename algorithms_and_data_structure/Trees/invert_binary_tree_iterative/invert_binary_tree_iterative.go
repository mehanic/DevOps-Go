package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func invertBinaryTreeIterative(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    stack := []*TreeNode{root}

    for len(stack) > 0 {
        // Pop the last node from stack
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // Swap left and right children
        node.Left, node.Right = node.Right, node.Left

        // Push children if not nil
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    return root
}
