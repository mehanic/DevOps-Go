package main

import (
    "fmt"
    "sort"
)

type KthLargest struct {
    k   int
    arr []int
}

func Constructor(k int, nums []int) KthLargest {
    sort.Ints(nums)
    return KthLargest{k: k, arr: nums}
}

func (this *KthLargest) Add(val int) int {
    this.arr = append(this.arr, val)
    sort.Ints(this.arr)
    return this.arr[len(this.arr)-this.k]
}

func main() {
    kth := Constructor(3, []int{4, 5, 8, 2})

    fmt.Println(kth.Add(3))  // → 4
    fmt.Println(kth.Add(5))  // → 5
    fmt.Println(kth.Add(10)) // → 5
    fmt.Println(kth.Add(9))  // → 8
    fmt.Println(kth.Add(4))  // → 8
}
