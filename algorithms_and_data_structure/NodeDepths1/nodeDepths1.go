package main

import "fmt"

type BinaryTree struct {
    Value int
    Left, Right *BinaryTree
}

// Основная функция
func NodeDepths(root *BinaryTree) int {
    return nodeDepthsHelper(root, 0)
}

// Вспомогательная рекурсивная функция
func nodeDepthsHelper(root *BinaryTree, depth int) int {
    if root == nil {
        return 0
    }
    // Суммируем текущую глубину и рекурсивно обрабатываем левое и правое поддеревья
    return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}

func main() {
    // Строим пример дерева
    root := &BinaryTree{Value: 1}
    root.Left = &BinaryTree{Value: 2}
    root.Right = &BinaryTree{Value: 3}
    root.Left.Left = &BinaryTree{Value: 4}
    root.Left.Right = &BinaryTree{Value: 5}

    // Получаем сумму всех глубин
    result := NodeDepths(root)
    fmt.Println("Сумма глубин всех узлов:", result) // Output: 6
}
