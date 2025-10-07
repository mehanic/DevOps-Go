package main

import (
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }
    list2.Next = mergeTwoLists(list1, list2.Next)
    return list2
}

func main() {
    // Створимо list1: 1 -> 3 -> 5
    l1 := &ListNode{Val: 1, Next: &ListNode{
        Val: 3, Next: &ListNode{
            Val: 5,
        },
    }}

    // Створимо list2: 2 -> 4 -> 6
    l2 := &ListNode{Val: 2, Next: &ListNode{
        Val: 4, Next: &ListNode{
            Val: 6,
        },
    }}

    // Зливаємо списки
    merged := mergeTwoLists(l1, l2)

    // Виводимо результат без printList
    for curr := merged; curr != nil; curr = curr.Next {
        fmt.Print(curr.Val)
        if curr.Next != nil {
            fmt.Print(" -> ")
        }
    }
    fmt.Println()
}
