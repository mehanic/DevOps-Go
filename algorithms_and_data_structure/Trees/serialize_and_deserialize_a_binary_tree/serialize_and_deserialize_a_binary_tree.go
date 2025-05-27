package main

import (
    "strconv"
    "strings"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func serialize(root *TreeNode) string {
    var serializedList []string
    preorderSerialize(root, &serializedList)
    return strings.Join(serializedList, ",")
}

func preorderSerialize(node *TreeNode, serializedList *[]string) {
    if node == nil {
        *serializedList = append(*serializedList, "#")
        return
    }
    *serializedList = append(*serializedList, strconv.Itoa(node.Val))
    preorderSerialize(node.Left, serializedList)
    preorderSerialize(node.Right, serializedList)
}

func deserialize(data string) *TreeNode {
    nodeValues := strings.Split(data, ",")
    index := 0
    return buildTree(nodeValues, &index)
}

func buildTree(values []string, index *int) *TreeNode {
    if *index >= len(values) {
        return nil
    }
    val := values[*index]
    *index++

    if val == "#" {
        return nil
    }

    intVal, _ := strconv.Atoi(val)
    node := &TreeNode{Val: intVal}
    node.Left = buildTree(values, index)
    node.Right = buildTree(values, index)

    return node
}
