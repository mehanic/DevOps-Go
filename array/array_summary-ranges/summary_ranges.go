package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[0]))
		return res
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			i += 1
		}
		if nums[i] != num {
			var tempBuffer bytes.Buffer
			tempBuffer.WriteString(strconv.Itoa(num))
			tempBuffer.WriteString("->")
			tempBuffer.WriteString(strconv.Itoa(nums[i]))
			res = append(res, tempBuffer.String())
		} else {
			res = append(res, strconv.Itoa(num))
		}
	}
	return res
}

func main() {
	// Test array
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("Input array:", nums)

	// Call summaryRanges function
	result := summaryRanges(nums)

	// Display the result
	fmt.Println("Summary of ranges:", result)
}

//Input array: [0 1 2 4 5 7]
//Summary of ranges: [0->2 4->5 7]

// Explanation of the Function
// Check for Single-Element Arrays:

// If the input array contains only one element, it appends that element directly to the result and returns since there’s nothing more to group.
// Iterate Through Elements:

// For each element num in the array, it identifies consecutive elements by checking if the next element is exactly one more than the current one.
// When it reaches the end of a sequence of consecutive numbers, it appends the range to the result.
// If the range contains more than one element, it uses bytes.Buffer to build a string in the format "start->end".
// If it’s a single number, it just appends that number.
// Return:

// The function returns the list of ranges in string format.
