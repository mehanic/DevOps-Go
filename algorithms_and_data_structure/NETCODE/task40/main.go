package main

import (
	"fmt"
	"sort"
)

//1. Brute Force
func singleNumber(nums []int) int {
    // Iterate through the array
    for i := 0; i < len(nums); i++ {
        flag := true
        // Compare the current element with all the other elements
        for j := 0; j < len(nums); j++ {
            // If a duplicate element is found, set flag to false and break
            if i != j && nums[i] == nums[j] {
                flag = false
                break
            }
        }
        // If flag is still true, it means the element is unique
        if flag {
            return nums[i]
        }
    }
    return 0 // Return 0 if no unique element is found (should not happen in this problem)
}

func main() {
    // Test case 1: Output should be 2 as it is the only unique number
    fmt.Println(singleNumber([]int{3, 2, 3}))       // Output: 2
    
    // Test case 2: Output should be 8 as it is the only unique number
    fmt.Println(singleNumber([]int{7, 6, 6, 7, 8})) // Output: 8
    
    // Test case 3: Output should be 5 as it is the only unique number
    fmt.Println(singleNumber([]int{1, 1, 2, 2, 5})) // Output: 5
    
    // Test case 4: Output should be 1 as it is the only unique number
    fmt.Println(singleNumber([]int{10, 20, 30, 10, 20})) // Output: 30
    
    // Test case 5: Output should be 9 as it is the only unique number
    fmt.Println(singleNumber([]int{5, 6, 7, 5, 6})) // Output: 7
}


//2. Hash Set

func singleNumber1(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			delete(seen, num)
		} else {
			seen[num] = true
		}
	}
	for num := range seen {
		return num
	}
	return -1
}

//3. Sorting

func singleNumber2(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			return nums[i]
		}
	}
	return nums[i]
}

//4. Bit Manipulation

func singleNumber3(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

