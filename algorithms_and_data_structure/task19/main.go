package main

import (
	"fmt"
)

// 1. Recursion
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive function to reverse the linked list
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next) // Reverse rest of the list
	head.Next.Next = head             // Reverse link
	head.Next = nil                   // Break old connection

	return newHead // Return new head
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

	fmt.Println("--------")
	main1()
}

//2. Iteration
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

// Main function to test the reverseList function
func main1() {
	// Example 1: [0,1,2,3]
	head := createList([]int{0, 4, 5, 6, 9})
	fmt.Print("Original List: ")
	printList(head)

	reversedHead := reverseList(head)
	fmt.Print("Reversed List: ")
	printList(reversedHead)

	// Example 2: []
	emptyHead := createList([]int{})
	fmt.Print("Original Empty List: ")
	printList(emptyHead)

	reversedEmpty := reverseList1(emptyHead)
	fmt.Print("Reversed Empty List: ")
	printList(reversedEmpty)
}
