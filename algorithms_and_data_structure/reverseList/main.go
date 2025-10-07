package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	if head.Next != nil {
		newHead = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil

	return newHead
}

func main() {
	// Створимо список 1 -> 2 -> 3
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	// Розвертаємо список
	reversed := reverseList(node1)

	// Виведемо результат
	for curr := reversed; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}


func main1() {
    node := &ListNode{Val: 42}
    reversed := reverseList(node)

    for curr := reversed; curr != nil; curr = curr.Next {
        fmt.Print(curr.Val)
    }
    fmt.Println()
}


func main2() {
    var node *ListNode = nil
    reversed := reverseList(node)

    for curr := reversed; curr != nil; curr = curr.Next {
        fmt.Print(curr.Val)
    }
    fmt.Println("Done") // Просто покажемо, що завершено
}

