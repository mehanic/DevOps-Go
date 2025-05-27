package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// linkedListLoop detects a cycle in the linked list using Floyd's Tortoise and Hare algorithm.
func linkedListLoop(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Example usage
func main() {
	// List with a cycle
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // creates a cycle

	fmt.Println(linkedListLoop(node1)) // Output: true

	// List without a cycle
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c

	fmt.Println(linkedListLoop(a)) // Output: false
}
