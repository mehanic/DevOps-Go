package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 5, 1, 8, 7, 9, 3, 0}
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Sort(sort.Reverse(sort.IntSlice(xi)))
	fmt.Println(xi)
}

// [0 1 3 3 5 7 8 9]
// [9 8 7 5 3 3 1 0]
