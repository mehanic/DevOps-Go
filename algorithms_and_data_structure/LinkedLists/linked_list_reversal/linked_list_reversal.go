package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListReversal(head *ListNode) *ListNode {
	var prevNode *ListNode = nil
	currNode := head

	// Переворачиваем направление указателей пока currNode не станет nil
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	// prevNode теперь указывает на новую голову перевёрнутого списка
	return prevNode
}
