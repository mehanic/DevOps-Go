package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Проверка, сбалансировано ли дерево
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    left := height(root.Left)
    right := height(root.Right)
    if abs(left-right) > 1 {
        return false
    }
    return isBalanced(root.Left) && isBalanced(root.Right)
}

// Вычисление высоты дерева
func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(height(root.Left), height(root.Right))
}

// Максимум из двух чисел
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Абсолютное значение
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    // Пример 1: сбалансированное дерево
    //
    //       1
    //      / \
    //     2   3
    //    / 
    //   4   
    root1 := &TreeNode{Val: 1}
    root1.Left = &TreeNode{Val: 2}
    root1.Right = &TreeNode{Val: 3}
    root1.Left.Left = &TreeNode{Val: 4}

    fmt.Println("Дерево 1 сбалансировано?", isBalanced(root1)) // true

    // Пример 2: несбалансированное дерево
    //
    //       1
    //      / 
    //     2   
    //    / 
    //   3
    root2 := &TreeNode{Val: 1}
    root2.Left = &TreeNode{Val: 2}
    root2.Left.Left = &TreeNode{Val: 3}

    fmt.Println("Дерево 2 сбалансировано?", isBalanced(root2)) // false
}
