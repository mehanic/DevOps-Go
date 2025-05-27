package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListReversalRecursive(head *ListNode) *ListNode {
	// Базовые случаи: пустой список или конец списка
	if head == nil || head.Next == nil {
		return head
	}

	// Рекурсивно переворачиваем оставшийся список
	newHead := linkedListReversalRecursive(head.Next)

	// Инвертируем текущую связь
	head.Next.Next = head
	head.Next = nil

	return newHead
}
