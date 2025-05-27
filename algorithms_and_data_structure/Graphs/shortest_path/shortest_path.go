package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	node, weight int
}

func shortestPath(n int, edges [][]int, start int) []int {
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt64
	}
	distances[start] = 0

	minHeap := &PriorityQueue{}
	heap.Init(minHeap)
	heap.Push(minHeap, &Item{node: start, dist: 0})

	for minHeap.Len() > 0 {
		curr := heap.Pop(minHeap).(*Item)
		currNode, currDist := curr.node, curr.dist

		if currDist > distances[currNode] {
			continue
		}

		for _, neighbor := range graph[currNode] {
			newDist := currDist + neighbor.weight
			if newDist < distances[neighbor.node] {
				distances[neighbor.node] = newDist
				heap.Push(minHeap, &Item{node: neighbor.node, dist: newDist})
			}
		}
	}

	// Replace unreachable distances with -1
	for i := range distances {
		if distances[i] == math.MaxInt64 {
			distances[i] = -1
		}
	}

	return distances
}

type Item struct {
	node int
	dist int
	index int // required by heap.Interface but unused here
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// Example usage
func main() {
	edges := [][]int{
		{0, 1, 4},
		{0, 2, 1},
		{2, 1, 2},
		{1, 3, 1},
		{2, 3, 5},
	}
	fmt.Println(shortestPath(4, edges, 0)) // Output: [0 3 1 4]
}
