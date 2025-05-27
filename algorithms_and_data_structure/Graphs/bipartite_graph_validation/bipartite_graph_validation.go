package main

func bipartiteGraphValidation(graph [][]int) bool {
    colors := make([]int, len(graph))

    for i := 0; i < len(graph); i++ {
        if colors[i] == 0 && !dfs(i, 1, graph, colors) {
            return false
        }
    }
    return true
}

func dfs(node int, color int, graph [][]int, colors []int) bool {
    colors[node] = color
    for _, neighbor := range graph[node] {
        if colors[neighbor] == color {
            return false
        }
        if colors[neighbor] == 0 && !dfs(neighbor, -color, graph, colors) {
            return false
        }
    }
    return true
}
