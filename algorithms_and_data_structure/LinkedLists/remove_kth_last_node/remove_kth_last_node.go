package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeKthLastNode(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	trailer := dummy
	leader := dummy

	// Advance 'leader' k steps ahead
	for i := 0; i < k; i++ {
		leader = leader.Next
		if leader == nil {
			// k is larger than the length of the list
			return head
		}
	}

	// Move both pointers until 'leader' reaches the end
	for leader.Next != nil {
		leader = leader.Next
		trailer = trailer.Next
	}

	// Remove the k-th node from the end
	trailer.Next = trailer.Next.Next

	return dummy.Next
}

func main() {
	// Создание списка: 1 -> 2 -> 3 -> 4 -> 5
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	// Удалить 2-й элемент с конца: результат 1 -> 2 -> 3 -> 5
	head := removeKthLastNode(n1, 2)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
