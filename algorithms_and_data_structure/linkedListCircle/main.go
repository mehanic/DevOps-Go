package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

// Функція перевірки циклу
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
    // Створюємо список без циклу: 1 -> 2 -> 3 -> nil
    node3 := &ListNode{Val: 3}
    node2 := &ListNode{Val: 2, Next: node3}
    node1 := &ListNode{Val: 1, Next: node2}

    fmt.Println("Has cycle (без циклу):", hasCycle(node1)) // false

    // Створюємо список з циклом: 1 -> 2 -> 3 -> 2 ...
    node3.Next = node2 // створюємо цикл
    fmt.Println("Has cycle (з циклом):", hasCycle(node1)) // true
}
