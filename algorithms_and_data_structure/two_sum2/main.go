package main

import "sort"

func twoSum(nums []int, target int) []int {
	// Map to store the indices of the numbers we have seen so far
	numberIndexMap := make(map[int]int)

	for currentIndex, currentNumber := range nums {
		// Calculate the number we need to reach the target
		numberNeeded := target - currentNumber

		// Check if the needed number exists in the map
		if neededIndex, found := numberIndexMap[numberNeeded]; found {
			return []int{neededIndex, currentIndex}
		}

		// Store the current number and its index in the map
		numberIndexMap[currentNumber] = currentIndex
	}

	// Return nil if no solution is found
	return nil
}

func twoSum1(nums []int, target int) []int {
    numMap := make(map[int]int) // A mapping to store numbers and their indices
    for i, num := range nums {
        complement := target - num // Find the required number to reach the target
        if index, found := numMap[complement]; found {
            return []int{index, i} // Return indices of the complement and current number
        }
        numMap[num] = i // Store the number with its index
    }
    return nil // This line is never reached due to the problem guarantee
}

func twoSum2(nums []int, target int) []int {
    for i := 1; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            if nums[j] + nums[j-i] == target {
                return []int{j-i, j}
            }
        }
    }
    return nil // Return nil if no solution found
}

func twoSum4(numbers []int, target int) []int {
    // Create a copy of the array with indices
    numbersWithIndices := make([][2]int, len(numbers))
    for i, num := range numbers {
        numbersWithIndices[i] = [2]int{num, i}
    }
    
    // Sort the array based on values
    sort.Slice(numbersWithIndices, func(i, j int) bool {
        return numbersWithIndices[i][0] < numbersWithIndices[j][0]
    })
    
    left, right := 0, len(numbers)-1
    
    for left < right {
        sum := numbersWithIndices[left][0] + numbersWithIndices[right][0]
        if sum == target {
            return []int{numbersWithIndices[left][1], numbersWithIndices[right][1]}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    
    return []int{} // Return empty slice if no solution found
}

//

func twoSum5(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        m[num] = i
    }
    for i, num := range nums {
        complement := target - num
        if j, ok := m[complement]; ok && j != i {
            return []int{i, j}
        }
    }
    return nil
}

//

func twoSum6(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := m[complement]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}

//

func twoSum7(nums []int, target int) []int {
    numMap := make(map[int]int) // A mapping to store numbers and their indices
    for i, num := range nums {
        complement := target - num // Find the required number to reach the target
        if index, found := numMap[complement]; found {
            return []int{index, i} // Return indices of the complement and current number
        }
        numMap[num] = i // Store the number with its index
    }
    return nil // This line is never reached due to the problem guarantee
}

func twoSum8(nums []int, target int) []int {
    var hash = make(map[int]int, 0)
    for index, num := range nums {
        if value, ok := hash[target - num]; ok {
            return []int{index, value}
        } else {
            hash[num] = index
        }
    }
    return nums
}

func twoSum9(nums []int, target int) []int {
    memo := map[int]int{}
    for i, num := range nums{
        if j, present := memo[target-num]; present {
            return []int{i,j}
        } else {
            memo[num] = i
        }
    }
    return []int{-1,-1}
}

func twoSum10(nums []int, target int) []int {
    numIndex:= make(map[int]int)
    for index,num := range nums{
        value:=target - num
        if j,found:= numIndex[value]; found{
            return []int{index,j}
        }
        numIndex[num]=index
    }
    return nil
}

func twoSum11(nums []int, target int) []int {
    seenMap := make(map[int]int)
    var diff int
    var result []int
    for i,v := range nums {
        diff = target - v
        if idx, ok := seenMap[diff]; ok {
            result = append(result, i, idx)
        }

        seenMap[v] = i 
    }
    return result
}

func twoSum12(nums []int, target int) []int {
    // Hash table to store number->index mapping
    ht := make(map[int]int)
    
    // Iterate through the array
    for i, num := range nums {
        // Check if complement exists in hash table
        if j, exists := ht[target-num]; exists {
            // If found, return indices of both numbers
            return []int{j, i}
        }
        
        // Store current number and its index in hash table
        ht[num] = i
    }
    
    // Return empty slice if no solution found
    return []int{}
}

// Time: O(n*n) = O(n^2)
// Space: O(1)
func twoSum13(nums []int, target int) []int {
    // Time: O(n)
    for i := 0; i < len(nums)-1; i++ {
        // Time: O(n)
        for j := i+1; j < len(nums); j++ {
			// Time: O(1)
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}

// Time: O(n+n) = O(2n) = O(n)
// Space: O(n)
func twoSum14(nums []int, target int) []int {
    // Space: O(n)
    m := make(map[int]int, 0)

    // Time: O(n)
    for idx, num := range nums {
        m[num] = idx
    }

    // Time: O(n)
    for indexX, num := range nums {
        // Time: O(1)
        if indexY, ok := m[target - num]; ok && indexX != indexY {
            return []int{indexX, indexY}
        }
    }

    return []int{}
}

// Time: O(n)
// Space: O(n)
func twoSum16(nums []int, target int) []int {
    // Space: O(n)
    s := make(map[int]int)

    // Time: O(n)
    for idx, num := range nums {
        // Time: O(1)
        if pos, ok := s[target-num]; ok {
            return []int{pos, idx}
        }
        s[num] = idx
    }
    return []int{}
}

// Time: O(n)
// Space: O(n)
func twoSum17(nums []int, target int) []int {
	// Space: O(n)
    s := newSet()
    
	// Time: O(n)
    for idx, num := range nums {
        if desired := target - num; s.has(desired) {
            idx2 := s.get(desired)
            return []int{idx2, idx}
        }
        s.add(num, idx)
    }
    return []int{}
}

type set struct {
    items map[int]int
}

func newSet() *set{
    return &set{
        items: make(map[int]int),
    }
}

func (s *set) add(value, position int) {
    s.items[value] = position
}

func (s *set) has(i int) bool {
    _, exists := s.items[i]
    return exists
}

func (s *set) get(i int) int {
    idx, exists := s.items[i]
    if !exists {
        return -1
    }
    return idx
}

// Time: O(n)
// Space: O(1)
// nums must be sorted!
func twoSum18(nums []int, target int) []int {
    if len(nums) <= 1 {
        return []int{}
    }

    // Time: O(n)
    var ptr1, ptr2 int = 0, len(nums)-1
    for ptr1 < ptr2 {
        current := nums[ptr2] + nums[ptr1]
		// Time: O(1) best O(3) worst so O(1)
        if current == target {
            return []int{ptr1, ptr2}
        } else if current > target {
            ptr2--
        } else /* if current < target */ {
            ptr1++
        }
    }
    
    return []int{}
}


