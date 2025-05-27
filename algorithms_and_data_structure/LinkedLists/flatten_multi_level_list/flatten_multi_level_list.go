package main

type MultiLevelListNode struct {
	Val   int
	Next  *MultiLevelListNode
	Child *MultiLevelListNode
}

func flattenMultiLevelList(head *MultiLevelListNode) *MultiLevelListNode {
	if head == nil {
		return nil
	}

	tail := head
	// Find the tail of the top-level list
	for tail.Next != nil {
		tail = tail.Next
	}

	curr := head
	for curr != nil {
		if curr.Child != nil {
			// Append child list to tail
			tail.Next = curr.Child
			curr.Child = nil

			// Update the tail to the new end of the list
			for tail.Next != nil {
				tail = tail.Next
			}
		}
		curr = curr.Next
	}

	return head
}
