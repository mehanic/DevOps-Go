package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"

	//s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	s.PrintBacking = true

	ages := []int{35, 15}
	s.Show("ages", ages)

	ages = append(ages, 5)
	s.Show("append(ages, 5)", ages)
	main1()
	main2()
	main3()
}

//---

func main1() {
	s.PrintBacking = true

	// #1: a nil slice has no backing array
	var nums []int
	s.Show("no backing array", nums)

	// #2: creates a new backing array
	nums = append(nums, 1, 3)
	s.Show("allocates", nums)

	// #3: creates a new backing array
	nums = append(nums, 2)
	s.Show("free capacity", nums)

	// #4: uses the same backing array
	nums = append(nums, 4)
	s.Show("no allocation", nums)

	// GOAL: append new odd numbers in the middle
	// [1 3 2 4] -> [1 3 7 9 2 4]

	// #5: [1 3 2 4] -> [1 3 2 4 2 4]
	nums = append(nums, nums[2:]...)
	s.Show("nums <- nums[2:]", nums)

	// #6: overwrites: [1 3 2 4 2 4] -> [1 3 7 9]
	nums = append(nums[:2], 7, 9)
	s.Show("nums[:2] <- 7, 9", nums)

	// #7: [1 3 7 9] -> [1 3 7 9 2 4]
	nums = nums[:6]
	s.Show("nums: extend", nums)
}

// don't mind about these options
// they're just for printing the slices nicely
func init() {
	s.MaxPerLine = 10
	s.Width = 45
}

//---

func main2() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}
}

//---

func main3() {
	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := float64(cap(ages))
		if c != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n",
				len(ages), c, c/oldCap)
		}
		oldCap = c
	}
}
