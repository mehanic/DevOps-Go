package main

import (
	"container/list"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreeColumns(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	columnMap := make(map[int][]int)
	leftmostColumn, rightmostColumn := 0, 0

	// Queue stores pairs of (*TreeNode, column index)
	queue := list.New()
	queue.PushBack(struct {
		node   *TreeNode
		column int
	}{root, 0})

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		curr := front.Value.(struct {
			node   *TreeNode
			column int
		})

		node, column := curr.node, curr.column
		if node != nil {
			columnMap[column] = append(columnMap[column], node.Val)

			if column < leftmostColumn {
				leftmostColumn = column
			}
			if column > rightmostColumn {
				rightmostColumn = column
			}

			queue.PushBack(struct {
				node   *TreeNode
				column int
			}{node.Left, column - 1})

			queue.PushBack(struct {
				node   *TreeNode
				column int
			}{node.Right, column + 1})
		}
	}

	result := make([][]int, 0, rightmostColumn-leftmostColumn+1)
	for i := leftmostColumn; i <= rightmostColumn; i++ {
		result = append(result, columnMap[i])
	}

	return result
}
