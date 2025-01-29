package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.MaxPerLine = 6
	s.PrintBacking = true

	ages := []int{35, 15, 25}
	s.Show("ages", ages)

	s.Show("ages[0:0]", ages[0:0])

	for i := 1; i < 4; i++ {
		txt := fmt.Sprintf("ages[%d:%d]", 0, i)
		s.Show(txt, ages[0:i])
	}

	s.Show("append", append(ages, 50))
	main1()
}

func main1() {
	s.PrintBacking = true

	// ----------------------------------------------------
	// #1 nil slice
	var games []string // nil slice
	s.Show("games", games)

	// ----------------------------------------------------
	// #2 empty slice
	games = []string{} // empty slice
	s.Show("games", games)
	// s.Show("another empty", []int{})

	// ----------------------------------------------------
	// #3 non-empty slice
	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)

	// ----------------------------------------------------
	// #4 reset the part using the games slice
	//    part is empty but its cap is still 4
	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part[:0]", part)
	s.Show("part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	// #6 backing array's elements become inaccessible
	// games = games[len(games):]

	// ----------------------------------------------------
	// #5 part doesn't have any more capacity
	//    games slice is still intact
	s.Show("part", part)
	s.Show("games", games)
}
