package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to detect a cycle in a linked list
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next        // Move slow pointer one step
		fast = fast.Next.Next   // Move fast pointer two steps

		if slow == fast { // If they meet, cycle exists
			return true
		}
	}

	return false // If fast reaches the end, no cycle
}

// Helper function to create a linked list from a slice and form a cycle
func createLinkedListWithCycle(values []int, cycleIndex int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	var cycleNode *ListNode

	for i, v := range values {
		newNode := &ListNode{Val: v}
		current.Next = newNode
		current = newNode

		if i == cycleIndex {
			cycleNode = newNode // Store the node where cycle should be
		}
	}

	// Create a cycle if cycleIndex is valid
	if cycleNode != nil {
		current.Next = cycleNode
	}

	return dummy.Next
}

func main() {
	// Example 1: List with cycle
	head1 := createLinkedListWithCycle([]int{1, 2, 3, 4}, 1)
	fmt.Println("Example 1:", hasCycle(head1)) // Output: true

	// Example 2: List without cycle
	head2 := createLinkedListWithCycle([]int{1, 2}, -1)
	fmt.Println("Example 2:", hasCycle(head2)) // Output: false

	// Example 3: Single node with cycle
	head3 := createLinkedListWithCycle([]int{5}, 0)
	fmt.Println("Example 3:", hasCycle(head3)) // Output: true

	// Example 4: Single node without cycle
	head4 := createLinkedListWithCycle([]int{10}, -1)
	fmt.Println("Example 4:", hasCycle(head4)) // Output: false
}

