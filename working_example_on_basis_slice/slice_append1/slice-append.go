package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	nums := []int{1, 2, 3}
	s.Show("nums", nums)

	_ = append(nums, 4)
	s.Show("nums", nums)

	nums = append(nums, 4)
	s.Show("nums", nums)

	nums = append(nums, 9)
	s.Show("nums", nums)

	nums = append(nums, 4)
	s.Show("nums", nums)

	// or:
	// nums = append(nums, 4, 9)
	// s.Show("nums", nums)

	nums = []int{1, 2, 3}
	tens := []int{12, 13}

	nums = append(nums, tens...)
	s.Show("nums", nums)
}

func main1() {
	var todo []string

	todo = append(todo, "sing")

	// you can only append elements with the same element type of the slice
	// todo = append(todo, 42)

	todo = append(todo, "run")

	// append is a variadic function, so you can append multiple elements
	todo = append(todo, "code", "play")

	// you can also append a slice to another slice using ellipsis: ...
	tomorrow := []string{"see mom", "learn go"}
	todo = append(todo, tomorrow...)
	// todo = append(todo, "see mom", "learn go")

	s.Show("todo", todo)
}


// nums                (len:3  cap:3  ptr:9616)
// ╔═══╗╔═══╗╔═══╗
// ║ 1 ║║ 2 ║║ 3 ║
// ╚═══╝╚═══╝╚═══╝
//   0    1    2  
//  nums                (len:3  cap:3  ptr:9616)
// ╔═══╗╔═══╗╔═══╗
// ║ 1 ║║ 2 ║║ 3 ║
// ╚═══╝╚═══╝╚═══╝
//   0    1    2  
//  nums                (len:4  cap:6  ptr:8096)
// ╔═══╗╔═══╗╔═══╗╔═══╗
// ║ 1 ║║ 2 ║║ 3 ║║ 4 ║
// ╚═══╝╚═══╝╚═══╝╚═══╝
//   0    1    2    3  
//  nums                (len:5  cap:6  ptr:8096)
// ╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗
// ║ 1 ║║ 2 ║║ 3 ║║ 4 ║║ 9 ║
// ╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝
//   0    1    2    3    4  
//  nums                (len:6  cap:6  ptr:8096)
// ╔═══╗╔═══╗╔═══╗╔═══╗╔═══╗
// ║ 1 ║║ 2 ║║ 3 ║║ 4 ║║ 9 ║
// ╚═══╝╚═══╝╚═══╝╚═══╝╚═══╝
//   0    1    2    3    4  
// ╔═══╗
// ║ 4 ║
// ╚═══╝
//   5  
//  nums                (len:5  cap:6  ptr:8288)
// ╔═══╗╔═══╗╔═══╗╔════╗╔════╗
// ║ 1 ║║ 2 ║║ 3 ║║ 12 ║║ 13 ║
// ╚═══╝╚═══╝╚═══╝╚════╝╚════╝
//   0    1    2     3     4  
