package main

import (
	"fmt"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

//2. Min-Heap
type KthLargest struct {
    minHeap *priorityqueue.Queue
    k       int
}

func Constructor(k int, nums []int) KthLargest {
    minHeap := priorityqueue.NewWith(utils.IntComparator)
    for _, num := range nums {
        minHeap.Enqueue(num)
    }
    for minHeap.Size() > k {
        minHeap.Dequeue()
    }
    return KthLargest{minHeap: minHeap, k: k}
}

func (this *KthLargest) Add(val int) int {
    this.minHeap.Enqueue(val)
    if this.minHeap.Size() > this.k {
        this.minHeap.Dequeue()
    }
    top, _ := this.minHeap.Peek()
    return top.(int)
}


func main() {
    // Initialize the KthLargest object with k = 3 and nums = [4, 5, 8, 2]
    kthLargest := Constructor(3, []int{4, 5, 8, 2})

    // Add 3 to the collection
    fmt.Println(kthLargest.Add(3))  // Expected output: 4 (3 is added, so 4th largest number becomes 4)

    // Add 5 to the collection
    fmt.Println(kthLargest.Add(5))  // Expected output: 5 (5 is added, so the 3rd largest number becomes 5)

    // Add 10 to the collection
    fmt.Println(kthLargest.Add(10)) // Expected output: 5 (10 is added, so the 3rd largest number remains 5)

    // Add 9 to the collection
    fmt.Println(kthLargest.Add(9))  // Expected output: 8 (9 is added, so the 3rd largest number becomes 8)

    // Add 4 to the collection
    fmt.Println(kthLargest.Add(4))  // Expected output: 8 (4 is added, so the 3rd largest number remains 8)
}