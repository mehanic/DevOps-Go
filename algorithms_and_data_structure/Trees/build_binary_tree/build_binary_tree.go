package main

// TreeNode definition
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

var preorderIndex int
var inorderIndexesMap map[int]int

func buildBinaryTree(preorder []int, inorder []int) *TreeNode {
    preorderIndex = 0
    inorderIndexesMap = make(map[int]int)

    // Populate the map with inorder values and their indices
    for i, val := range inorder {
        inorderIndexesMap[val] = i
    }

    return buildSubtree(0, len(inorder)-1, preorder)
}

func buildSubtree(left, right int, preorder []int) *TreeNode {
    if left > right {
        return nil
    }

    val := preorder[preorderIndex]
    inorderIndex := inorderIndexesMap[val]
    node := &TreeNode{Val: val}
    preorderIndex++

    node.Left = buildSubtree(left, inorderIndex-1, preorder)
    node.Right = buildSubtree(inorderIndex+1, right, preorder)

    return node
}
