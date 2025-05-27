package main

import (
	"fmt"
)

func prerequisites(n int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inDegrees := make([]int, n)

	// Build the graph and compute in-degrees
	for _, pair := range prerequisites {
		pre, course := pair[0], pair[1]
		graph[pre] = append(graph[pre], course)
		inDegrees[course]++
	}

	// Initialize queue with nodes having 0 in-degree
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	enrolledCourses := 0

	// Topological sort
	for len(queue) > 0 {
		// Dequeue
		node := queue[0]
		queue = queue[1:]
		enrolledCourses++

		for _, neighbor := range graph[node] {
			inDegrees[neighbor]--
			if inDegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we've visited all courses, return true
	return enrolledCourses == n
}

// Example usage
func main() {
	prereqs := [][]int{{1, 0}, {2, 1}, {3, 2}}
	fmt.Println(prerequisites(4, prereqs)) // Output: true
}
