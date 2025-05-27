package main

import "fmt"

// Определение ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Основная функция сортировки связанного списка
func sortLinkedList(head *ListNode) *ListNode {
	// Если список пуст или содержит один элемент, он уже отсортирован
	if head == nil || head.Next == nil {
		return head
	}

	// Разделить список на две половины
	secondHead := splitList(head)

	// Рекурсивно отсортировать обе половины
	firstSorted := sortLinkedList(head)
	secondSorted := sortLinkedList(secondHead)

	// Объединить отсортированные половины
	return merge(firstSorted, secondSorted)
}

// Функция разделения списка пополам с использованием двух указателей
func splitList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHead := slow.Next
	slow.Next = nil
	return secondHead
}

// Функция объединения двух отсортированных списков
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

// Пример для проверки
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Создание тестового связанного списка: 4 -> 2 -> 1 -> 3
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	sorted := sortLinkedList(head)
	printList(sorted) // Вывод: 1 2 3 4
}
