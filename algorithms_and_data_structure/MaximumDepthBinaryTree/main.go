package main

import (
    "fmt"
)

// Определение структуры узла дерева
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Функция вычисления максимальной глубины дерева
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// Вспомогательная функция для нахождения максимума двух чисел
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // Пример дерева:
    //       1
    //      / \
    //     2   3
    //    / 
    //   4
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}

    fmt.Println("Максимальная глубина дерева:", maxDepth(root))
}
