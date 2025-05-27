package main

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

var lca *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    lca = nil
    dfs(root, p, q)
    return lca
}

func dfs(node, p, q *TreeNode) bool {
    if node == nil {
        return false
    }

    nodeIsPorQ := node == p || node == q
    leftContains := dfs(node.Left, p, q)
    rightContains := dfs(node.Right, p, q)

    count := 0
    if nodeIsPorQ {
        count++
    }
    if leftContains {
        count++
    }
    if rightContains {
        count++
    }

    if count >= 2 {
        lca = node
    }

    return nodeIsPorQ || leftContains || rightContains
}
