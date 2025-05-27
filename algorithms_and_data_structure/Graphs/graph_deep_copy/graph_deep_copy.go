package main

// Definition of GraphNode
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func graphDeepCopy(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	cloneMap := make(map[*GraphNode]*GraphNode)
	return dfs(node, cloneMap)
}

func dfs(node *GraphNode, cloneMap map[*GraphNode]*GraphNode) *GraphNode {
	if val, exists := cloneMap[node]; exists {
		return val
	}
	clonedNode := &GraphNode{Val: node.Val}
	cloneMap[node] = clonedNode
	for _, neighbor := range node.Neighbors {
		clonedNeighbor := dfs(neighbor, cloneMap)
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
	}
	return clonedNode
}
