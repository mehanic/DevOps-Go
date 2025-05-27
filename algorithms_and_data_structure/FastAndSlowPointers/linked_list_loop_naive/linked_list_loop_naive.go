package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// linkedListLoopNaive detects a loop in a linked list using a map to store visited nodes.
func linkedListLoopNaive(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	curr := head
	for curr != nil {
		if visited[curr] {
			return true
		}
		visited[curr] = true
		curr = curr.Next
	}
	return false
}

// Helper function to create a simple test.
func main() {
	// Example with a loop
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // creates a loop

	fmt.Println(linkedListLoopNaive(node1)) // Output: true

	// Example without a loop
	nodeA := &ListNode{Val: 1}
	nodeB := &ListNode{Val: 2}
	nodeC := &ListNode{Val: 3}
	nodeA.Next = nodeB
	nodeB.Next = nodeC

	fmt.Println(linkedListLoopNaive(nodeA)) // Output: false
}
