/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func main() {// Точка входа

	// Создаём список: 1 -> 2 -> 3 -> 4 -> nil
	node4 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Print("Original list: ")
	printList(node1)

	// Разворачиваем список
	reversed := reverseList(node1)

	fmt.Print("Reversed list: ")
	printList(reversed)
}