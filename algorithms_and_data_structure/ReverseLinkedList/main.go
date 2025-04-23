package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse the linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next // Store next node
		curr.Next = prev  // Reverse the link
		prev = curr       // Move prev forward
		curr = next       // Move curr forward
	}

	return prev // New head of reversed list
}

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// Helper function to print a linked list
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}

// Main function to test the reverseList function
func main() {
	// Example 1: [0,1,2,3]
	head := createList([]int{0, 1, 2, 3})
	fmt.Print("Original List: ")
	printList(head)

	reversedHead := reverseList(head)
	fmt.Print("Reversed List: ")
	printList(reversedHead)

	// Example 2: []
	emptyHead := createList([]int{})
	fmt.Print("Original Empty List: ")
	printList(emptyHead)

	reversedEmpty := reverseList(emptyHead)
	fmt.Print("Reversed Empty List: ")
	printList(reversedEmpty)
}
