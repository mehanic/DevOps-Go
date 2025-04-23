package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. Recursion
// Function to merge two sorted linked lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// Helper function to create a linked list from a slice
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	list1 := createLinkedList([]int{1, 3, 5})
	list2 := createLinkedList([]int{2, 4, 6})

	fmt.Println("List 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList := mergeTwoLists(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 2: One empty list
	list1 = createLinkedList([]int{})
	list2 = createLinkedList([]int{0, 2, 3})

	fmt.Println("\nList 1 (empty):")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 3: Lists with one element each
	list1 = createLinkedList([]int{1})
	list2 = createLinkedList([]int{2})

	fmt.Println("\nList 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 4: Lists with duplicate values
	list1 = createLinkedList([]int{1, 2, 4})
	list2 = createLinkedList([]int{1, 3, 4})

	fmt.Println("\nList 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	fmt.Println("----------")
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
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	node.Next = list1
	if list1 == nil {
		node.Next = list2
	}

	return dummy.Next
}

func main1() {
	// Example 1
	list1 := createLinkedList([]int{1, 3, 5})
	list2 := createLinkedList([]int{2, 4, 6})

	fmt.Println("List 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList := mergeTwoLists1(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 2: One empty list
	list1 = createLinkedList([]int{})
	list2 = createLinkedList([]int{0, 2, 3})

	fmt.Println("\nList 1 (empty):")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists1(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 3: Lists with one element each
	list1 = createLinkedList([]int{1})
	list2 = createLinkedList([]int{2})

	fmt.Println("\nList 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists1(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)

	// Example 4: Lists with duplicate values
	list1 = createLinkedList([]int{1, 2, 4})
	list2 = createLinkedList([]int{1, 3, 4})

	fmt.Println("\nList 1:")
	printLinkedList(list1)
	fmt.Println("List 2:")
	printLinkedList(list2)

	mergedList = mergeTwoLists1(list1, list2)
	fmt.Println("Merged List:")
	printLinkedList(mergedList)
}
