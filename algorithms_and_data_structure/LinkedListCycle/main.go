/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if seen[cur] {
            return true
        }
        seen[cur] = true
        cur = cur.Next
    }
    return false
}


func main() {
	// No cycle: 1 -> 2 -> 3 -> nil
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Println("Has cycle (should be false):", hasCycle(node1))

	// Create a cycle: 1 -> 2 -> 3 -> 2 ...
	node3.Next = node2

	fmt.Println("Has cycle (should be true):", hasCycle(node1))
}