package main

import (
	"fmt"
)

// Interval struct represents a meeting time
type Interval struct {
	Start int
	End   int
}
//1. Brute Force
// canAttendMeetings checks if a person can attend all meetings without conflict
func canAttendMeetings(intervals []Interval) bool {
	n := len(intervals)
	for i := 0; i < n; i++ {
		A := intervals[i]
		for j := i + 1; j < n; j++ {
			B := intervals[j]
			if min(A.End, B.End) > max(A.Start, B.Start) {
				return false
			}
		}
	}
	return true
}

// Utility function to find the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Utility function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1: Overlapping meetings
	intervals1 := []Interval{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println("Example 1:", canAttendMeetings(intervals1)) // Output: false

	// Example 2: No overlapping meetings
	intervals2 := []Interval{{5, 8}, {9, 15}}
	fmt.Println("Example 2:", canAttendMeetings(intervals2)) // Output: true

	// Example 3: Edge case - No meetings
	intervals3 := []Interval{}
	fmt.Println("Example 3:", canAttendMeetings(intervals3)) // Output: true

	// Example 4: Meetings that just touch
	intervals4 := []Interval{{0, 5}, {5, 10}, {10, 15}}
	fmt.Println("Example 4:", canAttendMeetings(intervals4)) // Output: true
}


//2. Sorting

/**
* Definition of Interval:
* type Interval struct {
*    start int
*    end   int
* }
*/

// func canAttendMeetings2(intervals []Interval) bool {
//     sort.Slice(intervals, func(i, j int) bool {
//         return intervals[i].start < intervals[j].start
//     })

//     for i := 1; i < len(intervals); i++ {
//         if intervals[i-1].end > intervals[i].start {
//             return false
//         }
//     }
//     return true
// }