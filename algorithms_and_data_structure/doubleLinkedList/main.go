package main

import "fmt"

// Node structure for Doubly Linked List
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

// DoublyLinkedList structure
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// Constructor to create a new Doubly Linked List
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// O(1) time | O(1) space
// SetHead sets the given node as the head of the list
func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

// O(1) time | O(1) space
// SetTail sets the given node as the tail of the list
func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

// O(1) time | O(1) space
// InsertBefore inserts a node before a given node
func (ll *DoublyLinkedList) InsertBefore(existingNode *Node, newNode *Node) {
	if existingNode == nil {
		return
	}
	// If inserting before head, update head
	if existingNode == ll.Head {
		ll.Head = newNode
	}

	newNode.Prev = existingNode.Prev
	newNode.Next = existingNode
	if existingNode.Prev != nil {
		existingNode.Prev.Next = newNode
	}
	existingNode.Prev = newNode
}

// O(1) time | O(1) space
// InsertAfter inserts a node after a given node
func (ll *DoublyLinkedList) InsertAfter(existingNode *Node, newNode *Node) {
	if existingNode == nil {
		return
	}
	// If inserting after tail, update tail
	if existingNode == ll.Tail {
		ll.Tail = newNode
	}

	newNode.Next = existingNode.Next
	newNode.Prev = existingNode
	if existingNode.Next != nil {
		existingNode.Next.Prev = newNode
	}
	existingNode.Next = newNode
}

// O(n) time | O(1) space
// PrintList prints the list from head to tail
func (ll *DoublyLinkedList) PrintList() {
	curr := ll.Head
	for curr != nil {
		fmt.Print(curr.Value, " <-> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}

// O(n) time | O(1) space
// PrintReverse prints the list from tail to head
func (ll *DoublyLinkedList) PrintReverse() {
	curr := ll.Tail
	for curr != nil {
		fmt.Print(curr.Value, " <-> ")
		curr = curr.Prev
	}
	fmt.Println("nil")
}

func main() {
	// Creating a new Doubly Linked List
	dll := NewDoublyLinkedList()

	// Creating nodes
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	// Setting Head and Tail
	dll.SetHead(node1)
	dll.SetTail(node4)

	// Inserting nodes
	dll.InsertAfter(node1, node2) // 1 <-> 2
	dll.InsertBefore(node4, node3) // 1 <-> 2 <-> 3 <-> 4

	// Printing the list
	fmt.Println("Forward:")
	dll.PrintList() // Output: 1 <-> 2 <-> 3 <-> 4 <-> nil

	fmt.Println("Reverse:")
	dll.PrintReverse() // Output: 4 <-> 3 <-> 2 <-> 1 <-> nil
}
