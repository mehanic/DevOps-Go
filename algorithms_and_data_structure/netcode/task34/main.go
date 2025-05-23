package main

import (
	"fmt"
	"sort"
	"github.com/emirpasic/gods/queues/priorityqueue"
)
//1. Sorting
func lastStoneWeight(stones []int) int {
	// Continue until there is one or fewer stones left
    for len(stones) > 1 {
		// Sort the stones in ascending order to get the heaviest two stones
		sort.Ints(stones)
		// Subtract the two heaviest stones
		cur := stones[len(stones)-1] - stones[len(stones)-2]
		// Remove the two heaviest stones from the list
		stones = stones[:len(stones)-2]
		// If there's a positive difference, add the new stone to the list
		if cur > 0 {
			stones = append(stones, cur)
		}
	}
	// If no stones are left, return 0; otherwise, return the remaining stone
	if len(stones) == 0 {
		return 0
	}
	return stones[0]   
}

func main() {
    // Example 1
	stones1 := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones1)) // Output: 1 (Explanation below)

    // Example 2
	stones2 := []int{10, 4, 2, 10}
	fmt.Println(lastStoneWeight(stones2)) // Output: 4 (Explanation below)

    // Example 3
	stones3 := []int{5, 3, 2, 1}
	fmt.Println(lastStoneWeight(stones3)) // Output: 1 (Explanation below)
}


//2. Binary Search

func lastStoneWeight1(stones []int) int {
    sort.Ints(stones)
    n := len(stones)
    
    for n > 1 {
        cur := stones[n-1] - stones[n-2]
        n -= 2
        
        if cur > 0 {
            pos := sort.Search(n, func(i int) bool {
                return stones[i] >= cur
            })
            
            for i := n; i > pos; i-- {
                stones[i] = stones[i-1]
            }
            stones[pos] = cur
            n++
        }
    }
    
    if n > 0 {
        return stones[0]
    }
    return 0
}


//3. Heap

func lastStoneWeight2(stones []int) int {
    pq := priorityqueue.NewWith(func(a, b interface{}) int {
        return a.(int) - b.(int)  
    })
    
    for _, s := range stones {
        pq.Enqueue(-s)
    }
    
    for pq.Size() > 1 {
        first, _ := pq.Dequeue()
        second, _ := pq.Dequeue()
        if second.(int) > first.(int) {
            pq.Enqueue(first.(int) - second.(int))
        }
    }
    
    pq.Enqueue(0)
    result, _ := pq.Dequeue()
    return -result.(int)
}


//4. Bucket Sort

func lastStoneWeight3(stones []int) int {
    maxStone := 0
    for _, stone := range stones {
        if stone > maxStone {
            maxStone = stone
        }
    }
    
    bucket := make([]int, maxStone+1)
    for _, stone := range stones {
        bucket[stone]++
    }
    
    first, second := maxStone, maxStone
    for first > 0 {
        if bucket[first]%2 == 0 {
            first--
            continue
        }
        
        j := min(first-1, second)
        for j > 0 && bucket[j] == 0 {
            j--
        }
        
        if j == 0 {
            return first
        }
        second = j
        bucket[first]--
        bucket[second]--
        bucket[first-second]++
        first = max(first-second, second)
    }
    return first
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}