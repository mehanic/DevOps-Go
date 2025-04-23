//1. Hash Set

package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to detect if a cycle exists in a linked list
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if seen[cur] {
			return true
		}
		seen[cur] = true
		cur = cur.Next
	}
	return false
}

// Helper function to create a linked list from a slice (without cycle)
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

// Helper function to print a linked list (stops at 10 nodes to avoid infinite loops)
func printList(head *ListNode) {
	seen := make(map[*ListNode]bool)
	count := 0
	for head != nil {
		if seen[head] || count > 10 { // Prevent infinite loop in case of a cycle
			fmt.Print("... (cycle detected)")
			break
		}
		fmt.Print(head.Val, " -> ")
		seen[head] = true
		head = head.Next
		count++
	}
	fmt.Println("nil")
}

// Function to create a linked list with a cycle
func createCycleList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	var cycleNode *ListNode

	// Create nodes and find the position for cycle
	for i, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
		if i == pos-1 { // Store the cycle entry point
			cycleNode = curr
		}
	}

	// Create the cycle if pos is valid
	if cycleNode != nil {
		curr.Next = cycleNode
	}

	return head
}

func main() {
	// Example 1: No cycle
	list1 := createList([]int{1, 2, 3, 4})
	fmt.Print("List1: ")
	printList(list1)
	fmt.Println("Has Cycle?", hasCycle(list1)) // Expected: false

	// Example 2: Cycle present (tail connects to node at index 1)
	list2 := createCycleList([]int{1, 2, 3, 4, 5}, 1)
	fmt.Print("\nList2: ")
	printList(list2)
	fmt.Println("Has Cycle?", hasCycle(list2)) // Expected: true

	// Example 3: Cycle present (tail connects to node at index 3)
	list3 := createCycleList([]int{10, 20, 30, 40, 50, 60}, 3)
	fmt.Print("\nList3: ")
	printList(list3)
	fmt.Println("Has Cycle?", hasCycle(list3)) // Expected: true

	// Example 4: Single node, no cycle
	list4 := createList([]int{100})
	fmt.Print("\nList4: ")
	printList(list4)
	fmt.Println("Has Cycle?", hasCycle(list4)) // Expected: false

	// Example 5: Empty list
	list5 := createList([]int{})
	fmt.Print("\nList5: ")
	printList(list5)
	fmt.Println("Has Cycle?", hasCycle(list5)) // Expected: false
	fmt.Println("---------------")
	main1()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle1(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main1() {
	// Example 1: No cycle
	list1 := createList([]int{1, 2, 3, 4})
	fmt.Print("List1: ")
	printList(list1)
	fmt.Println("Has Cycle?", hasCycle1(list1)) // Expected: false

	// Example 2: Cycle present (tail connects to node at index 1)
	list2 := createCycleList([]int{1, 2, 3, 4, 5}, 1)
	fmt.Print("\nList2: ")
	printList(list2)
	fmt.Println("Has Cycle?", hasCycle1(list2)) // Expected: true

	// Example 3: Cycle present (tail connects to node at index 3)
	list3 := createCycleList([]int{10, 20, 30, 40, 50, 60}, 3)
	fmt.Print("\nList3: ")
	printList(list3)
	fmt.Println("Has Cycle?", hasCycle1(list3)) // Expected: true

	// Example 4: Single node, no cycle
	list4 := createList([]int{100})
	fmt.Print("\nList4: ")
	printList(list4)
	fmt.Println("Has Cycle?", hasCycle1(list4)) // Expected: false

	// Example 5: Empty list
	list5 := createList([]int{})
	fmt.Print("\nList5: ")
	printList(list5)
	fmt.Println("Has Cycle?", hasCycle1(list5)) // Expected: false
}
