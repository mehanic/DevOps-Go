package main

import "fmt"

// ListNode определяет узел односвязного списка.
type ListNode struct {
	Val  int
	Next *ListNode
}

// linkedListMidpoint возвращает узел, находящийся в середине списка.
func linkedListMidpoint(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Пример использования
func main() {
	// Создаем связный список: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	mid := linkedListMidpoint(head)
	fmt.Println(mid.Val) // Вывод: 3
}
