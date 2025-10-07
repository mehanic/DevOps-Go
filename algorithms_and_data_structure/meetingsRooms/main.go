package main

import (
    "fmt"
)

// Определение интервала
type Interval struct {
    start int
    end   int
}

func canAttendMeetings(intervals []Interval) bool {
    n := len(intervals)
    for i := 0; i < n; i++ {
        A := intervals[i]
        for j := i + 1; j < n; j++ {
            B := intervals[j]
            if min(A.end, B.end) > max(A.start, B.start) {
                return false
            }
        }
    }
    return true
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    meetings1 := []Interval{
        {start: 0, end: 30},
        {start: 5, end: 10},
        {start: 15, end: 20},
    }

    meetings2 := []Interval{
        {start: 5, end: 8},
        {start: 10, end: 15},
    }

    fmt.Println("Meetings1 - Can attend all?", canAttendMeetings(meetings1)) // false
    fmt.Println("Meetings2 - Can attend all?", canAttendMeetings(meetings2)) // true
}
