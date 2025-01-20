package main

import "fmt"

func main() {
	fruits := map[int]string{1: "mango", 2: "Orange"}
	fmt.Println(fruits)

	students := map[string][]string{
		"ss1": {"John", "Ade"},
		"ss2": {"John", "Ade"},
	}

	ss1Students := students["ss1"]

	fmt.Println(ss1Students)

	fmt.Println(students)
}

// map[1:mango 2:Orange]
// [John Ade]
// map[ss1:[John Ade] ss2:[John Ade]]
