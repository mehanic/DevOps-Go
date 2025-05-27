package main

import "container/list"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func rightmostNodesOfBinaryTree(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var res []int
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        levelSize := queue.Len()
        var node *TreeNode

        for i := 0; i < levelSize; i++ {
            front := queue.Front()
            queue.Remove(front)
            node = front.Value.(*TreeNode)

            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        // After the inner loop, 'node' points to the last node of this level
        res = append(res, node.Val)
    }

    return res
}
