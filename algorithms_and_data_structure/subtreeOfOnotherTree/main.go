package main

import (
    "fmt"
)

// Определение узла дерева
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Проверка, является ли subRoot поддеревом root
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true
    }
    if root == nil {
        return false
    }

    if sameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Проверка, идентичны ли два дерева
func sameTree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    if root != nil && subRoot != nil && root.Val == subRoot.Val {
        return sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right)
    }
    return false
}

func main() {
    // Главное дерево:
    //       3
    //      / \
    //     4   5
    //    / \
    //   1   2
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 5}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 2}

    // Поддерево:
    //     4
    //    / \
    //   1   2
    subRoot := &TreeNode{Val: 4}
    subRoot.Left = &TreeNode{Val: 1}
    subRoot.Right = &TreeNode{Val: 2}

    result := isSubtree(root, subRoot)
    fmt.Println("Является ли subRoot поддеревом root?", result) // true
}
