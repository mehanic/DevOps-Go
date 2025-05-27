package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func kthSmallestNumberInBSTRecursive(root *TreeNode, k int) int {
    sortedList := inorder(root)
    if k-1 < 0 || k-1 >= len(sortedList) {
        return -1 // handle invalid k
    }
    return sortedList[k-1]
}

// inorder traversal returns a sorted slice of node values from BST
func inorder(node *TreeNode) []int {
    if node == nil {
        return []int{}
    }
    left := inorder(node.Left)
    right := inorder(node.Right)
    return append(append(left, node.Val), right...)
}
