package main

import (
	"fmt"
	"math"
	"sort"
)

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}
	if uf.size[rootX] > uf.size[rootY] {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	} else {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	}
	return true
}

func connectTheDots(points [][]int) int {
	n := len(points)
	edges := make([][3]int, 0)

	// Build edge list with Manhattan distance as weight
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cost := int(math.Abs(float64(points[i][0]-points[j][0])) +
				math.Abs(float64(points[i][1]-points[j][1])))
			edges = append(edges, [3]int{cost, i, j})
		}
	}

	// Sort edges by cost
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	uf := NewUnionFind(n)
	totalCost := 0
	edgesAdded := 0

	// Kruskal's algorithm
	for _, edge := range edges {
		cost, u, v := edge[0], edge[1], edge[2]
		if uf.Union(u, v) {
			totalCost += cost
			edgesAdded++
			if edgesAdded == n-1 {
				break
			}
		}
	}

	return totalCost
}

// Example usage
func main() {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(connectTheDots(points)) // Output: 20
}
