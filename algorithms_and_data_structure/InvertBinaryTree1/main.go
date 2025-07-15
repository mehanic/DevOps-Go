/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := arrayqueue.New()
    queue.Enqueue(root)

    for queue.Size() > 0 {
        node, _ := queue.Dequeue()
        current := node.(*TreeNode)
        current.Left, current.Right = current.Right, current.Left

        if current.Left != nil {
            queue.Enqueue(current.Left)
        }
        if current.Right != nil {
            queue.Enqueue(current.Right)
        }
    }
    return root
}