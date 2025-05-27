package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := findMiddle(head)
	secondHead := reverseList(mid)

	ptr1 := head
	ptr2 := secondHead
	for ptr2 != nil {
		if ptr1.Val != ptr2.Val {
			return false
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
