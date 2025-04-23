package main

import (
	"fmt"
)

// Определение структуры узла
type Node struct {
	Name     string
	Children []*Node
}

// Метод для выполнения обхода в глубину (DFS)
func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name) // Добавляем текущий узел в массив
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array) // Рекурсивно вызываем DFS для каждого ребенка
	}
	return array
}

// Функция для создания тестового дерева и вызова DFS
func main() {
	// Создаем дерево
	root := &Node{Name: "A"}
	root.Children = []*Node{
		{Name: "B", Children: []*Node{
			{Name: "E"},
			{Name: "F"},
		}},
		{Name: "C"},
		{Name: "D", Children: []*Node{
			{Name: "G"},
			{Name: "H"},
		}},
	}

	// Запускаем DFS
	result := root.DepthFirstSearch([]string{})

	// Выводим результат обхода
	fmt.Println("Depth-First Search Result:", result)

	fmt.Println("------------")
	main1()
}

func main1() {
	// Creating a sample tree
	root := &Node{Name: "A"}
	root.Children = []*Node{
		{Name: "B", Children: []*Node{
			{Name: "E"},
			{Name: "F"},
		}},
		{Name: "C"},
		{Name: "D", Children: []*Node{
			{Name: "G"},
			{Name: "H"},
		}},
	}

	// Performing DFS traversal
	result := root.DepthFirstSearch([]string{})

	// Printing the result
	fmt.Println("DFS Traversal:", result) // Expected Output: [A B E F C D G H]
}
