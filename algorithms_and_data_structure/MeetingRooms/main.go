package main

import (
	"fmt"
	"sort"
)

// Interval struct represents a meeting time
type Interval struct {
	Start int
	End   int
}

// canAttendMeetings checks if a person can attend all meetings without conflict
func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// Check for overlaps
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}

	return true
}

func main() {
	// Example 1: Overlapping meetings
	intervals1 := []Interval{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(canAttendMeetings(intervals1)) // Output: false

	// Example 2: No overlapping meetings
	intervals2 := []Interval{{5, 8}, {9, 15}}
	fmt.Println(canAttendMeetings(intervals2)) // Output: true

	// Example 3: Edge case - No meetings
	intervals3 := []Interval{}
	fmt.Println(canAttendMeetings(intervals3)) // Output: true

	// Example 4: Meetings that just touch
	intervals4 := []Interval{{0, 5}, {5, 10}, {10, 15}}
	fmt.Println(canAttendMeetings(intervals4)) // Output: true
}
