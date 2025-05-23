package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	sliceable := []byte{'f', 'u', 'l', 'l'}

	s.Show("sliceable", sliceable)
	s.Show("sliceable[0:3]", sliceable[0:3])
	s.Show("sliceable[0:3:3]", sliceable[0:3:3])
	s.Show("sliceable[0:2:2]", sliceable[0:2:2])
	s.Show("sliceable[0:1:1]", sliceable[0:1:1])
	s.Show("sliceable[1:3:3]", sliceable[1:3:3])
	s.Show("sliceable[2:3:3]", sliceable[2:3:3])
	s.Show("sliceable[2:3:4]", sliceable[2:3:4])
	s.Show("sliceable[4:4:4]", sliceable[4:4:4])
}

//---

func main1() {
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4} // #1
	// odds := nums[:2]          // #2
	// odds := nums[:2:2]        // #4
	// odds = append(odds, 5, 7) // #3

	// odds := append(nums[:2:2], 5, 7) // #5
	// evens := append(nums[2:4], 6, 8) // #6

	s.Show("nums", nums) // #1
	// s.Show("odds", odds)   // #2
	// s.Show("evens", evens) // #6
}

// don't mind about these options
// they're just for printing the slices nicely
func init() {
	s.MaxPerLine = 10
	s.Width = 55
}
