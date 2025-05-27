package main

import (
    "container/list"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func widestBinaryTreeLevel(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxWidth := 0
    // Use a queue to hold pairs of (node, index)
    queue := list.New()
    queue.PushBack(struct {
        node  *TreeNode
        index int
    }{root, 0})

    for queue.Len() > 0 {
        levelSize := queue.Len()
        leftmostIndex := queue.Front().Value.(struct {
            node  *TreeNode
            index int
        }).index
        rightmostIndex := leftmostIndex
        
        for i := 0; i < levelSize; i++ {
            front := queue.Remove(queue.Front()).(struct {
                node  *TreeNode
                index int
            })
            node, idx := front.node, front.index
            
            if node.Left != nil {
                queue.PushBack(struct {
                    node  *TreeNode
                    index int
                }{node.Left, 2*idx + 1})
            }
            if node.Right != nil {
                queue.PushBack(struct {
                    node  *TreeNode
                    index int
                }{node.Right, 2*idx + 2})
            }
            rightmostIndex = idx
        }
        width := rightmostIndex - leftmostIndex + 1
        if width > maxWidth {
            maxWidth = width
        }
    }
    
    return maxWidth
}
