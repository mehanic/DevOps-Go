package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Функция проверки, одинаковы ли два дерева
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p != nil && q != nil && p.Val == q.Val {
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    return false
}

func main() {
    // Создаем первое дерево
    //
    //       1
    //      / \
    //     2   3
    root1 := &TreeNode{Val: 1}
    root1.Left = &TreeNode{Val: 2}
    root1.Right = &TreeNode{Val: 3}

    // Создаем второе дерево, идентичное первому
    root2 := &TreeNode{Val: 1}
    root2.Left = &TreeNode{Val: 2}
    root2.Right = &TreeNode{Val: 3}

    // Создаем третье дерево, отличающееся от первых двух
    //
    //       1
    //      / \
    //     3   2
    root3 := &TreeNode{Val: 1}
    root3.Left = &TreeNode{Val: 3}
    root3.Right = &TreeNode{Val: 2}

    fmt.Println("root1 и root2 одинаковы?", isSameTree(root1, root2)) // true
    fmt.Println("root1 и root3 одинаковы?", isSameTree(root1, root3)) // false
}
