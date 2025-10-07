package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Функция инверсии дерева
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        // Меняем левое и правое поддерево
        current.Left, current.Right = current.Right, current.Left

        if current.Left != nil {
            queue = append(queue, current.Left)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }

    return root
}

// Функция для печати дерева (BFS)
func printTree(root *TreeNode) {
    if root == nil {
        fmt.Println("nil")
        return
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current != nil {
            fmt.Print(current.Val, " ")
            queue = append(queue, current.Left, current.Right)
        } else {
            fmt.Print("nil ")
        }
    }
    fmt.Println()
}

func main() {
    // Создаём дерево
    root := &TreeNode{Val: 4}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 7}
    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Left = &TreeNode{Val: 6}
    root.Right.Right = &TreeNode{Val: 9}

    fmt.Println("Before inversion:")
    printTree(root)

    inverted := invertTree(root)

    fmt.Println("After inversion:")
    printTree(inverted)
}
