package main

import (
	"fmt"
	"sort"
)

// Definition of Interval struct
type Interval struct {
	start int
	end   int
}

// Function to check if a person can attend all meetings without conflict
func canAttendMeetings2(intervals []Interval) bool {
	// Sorting intervals by their start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	// Check for conflicts by comparing each pair of consecutive intervals
	for i := 1; i < len(intervals); i++ {
		// If the previous meeting ends after the current one starts, there is a conflict
		if intervals[i-1].end > intervals[i].start {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1: Meetings with no conflict
	intervals1 := []Interval{
		{start: 0, end: 30},
		{start: 35, end: 45},
		{start: 50, end: 60},
	}
	fmt.Println("Can attend all meetings (Test 1):", canAttendMeetings2(intervals1)) // Output: true

	// Test case 2: Meetings with conflict
	intervals2 := []Interval{
		{start: 0, end: 30},
		{start: 25, end: 35},
		{start: 40, end: 50},
	}
	fmt.Println("Can attend all meetings (Test 2):", canAttendMeetings2(intervals2)) // Output: false

	// Test case 3: Single meeting (no conflict)
	intervals3 := []Interval{
		{start: 10, end: 20},
	}
	fmt.Println("Can attend all meetings (Test 3):", canAttendMeetings2(intervals3)) // Output: true

	// Test case 4: Multiple meetings with one conflict
	intervals4 := []Interval{
		{start: 0, end: 30},
		{start: 10, end: 20},
		{start: 35, end: 45},
	}
	fmt.Println("Can attend all meetings (Test 4):", canAttendMeetings2(intervals4)) // Output: false
}

//

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
