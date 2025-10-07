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

// Функция подсчета диаметра дерева
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

// Вспомогательная функция — высота дерева
func maxHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(maxHeight(root.Left), maxHeight(root.Right))
}

// Максимум из двух чисел
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Пример дерева:
    //        1
    //       / \
    //      2   3
    //     / \     
    //    4   5    
    //
    // Диаметр — путь [4 -> 2 -> 1 -> 3] (3 ребра)

    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}

    fmt.Println("Диаметр дерева:", diameterOfBinaryTree(root))
}
