package main

import (
	"strings"

	prettyslice "github.com/inancgumus/prettyslice"
	s "github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true

	prettyslice.Show("make([]int, 3)", make([]int, 3))
	prettyslice.Show("make([]int, 3, 5)", make([]int, 3, 5))

	s := make([]int, 0, 5)
	prettyslice.Show("make([]int, 0, 5)", s)

	s = append(s, 42)
	prettyslice.Show("s = append(s, 42)", s)
}

func main1() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	// #1: assume that tasks can be a long list
	// ---------------------------------------------
	tasks := []string{"jump", "run", "read"}

	// #2: INEFFICIENT WAY
	// ---------------------------------------------
	// var upTasks []string
	// s.Show("upTasks", upTasks)

	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }

	// #3: SO SO WAY:
	// ---------------------------------------------
	// upTasks := make([]string, len(tasks))
	// s.Show("upTasks", upTasks)

	// for _, task := range tasks {
	// 	upTasks = append(upTasks, strings.ToUpper(task))
	// 	s.Show("upTasks", upTasks)
	// }

	// #4: SO SO WAY 2:
	// ---------------------------------------------
	// upTasks := make([]string, len(tasks))
	// s.Show("upTasks", upTasks)

	// for i, task := range tasks {
	// 	upTasks[i] = strings.ToUpper(task)
	// 	s.Show("upTasks", upTasks)
	// }

	// #5: allocates a new array
	// upTasks = append(upTasks, "PLAY")
	// s.Show("upTasks", upTasks)

	// #6: SO SO WAY 3
	// ---------------------------------------------
	// upTasks := make([]string, 0, len(tasks))

	// #7
	// upTasks = upTasks[:cap(upTasks)]

	// #6
	// s.Show("upTasks", upTasks)

	// for i, task := range tasks {
	// 	upTasks[i] = strings.ToUpper(task)
	// 	s.Show("upTasks", upTasks)
	// }

	// #8: THE BEST WAY
	// ---------------------------------------------
	upTasks := make([]string, 0, len(tasks))

	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}
}
