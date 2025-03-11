package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {

	xs := people{"Ben", "Apple", "Banana", "No", "Kaaa"}

	// asc
	sort.Strings(xs)
	fmt.Println(xs)

	// desc
	sort.Sort(sort.Reverse(sort.StringSlice(xs)))
	fmt.Println(xs)
}
