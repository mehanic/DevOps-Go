package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListIntersection(headA, headB *ListNode) *ListNode {
	ptrA := headA
	ptrB := headB

	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}

		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}

	return ptrA // Could be the intersection node or nil
}
