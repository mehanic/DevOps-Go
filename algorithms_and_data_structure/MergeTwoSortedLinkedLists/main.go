package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive function to merge two sorted linked lists
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
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

// Main function to test mergeTwoLists
func main() {
	// Example 1: [1,2,4] and [1,3,5]
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 5})

	fmt.Print("List1: ")
	printList(list1)
	fmt.Print("List2: ")
	printList(list2)

	mergedList := mergeTwoLists(list1, list2)
	fmt.Print("Merged List: ")
	printList(mergedList)

	// Example 2: [] and [1,2]
	list3 := createList([]int{})
	list4 := createList([]int{1, 2})

	fmt.Print("\nList3: ")
	printList(list3)
	fmt.Print("List4: ")
	printList(list4)

	mergedList2 := mergeTwoLists(list3, list4)
	fmt.Print("Merged List: ")
	printList(mergedList2)

	// Example 3: [] and []
	list5 := createList([]int{})
	list6 := createList([]int{})

	fmt.Print("\nList5: ")
	printList(list5)
	fmt.Print("List6: ")
	printList(list6)

	mergedList3 := mergeTwoLists(list5, list6)
	fmt.Print("Merged List: ")
	printList(mergedList3)
}
