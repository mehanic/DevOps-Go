package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// with arrays
func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5
	var uniques [max]int

	// It's harder to make a program dynamic using arrays
	// max, _ := strconv.Atoi(os.Args[1])
	// var uniques [10]int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques[found] = n
		found++
	}

	fmt.Println("\n\nuniques:", uniques)
	main1()
}

//--

// with slices
func main1() {
	rand.Seed(time.Now().UnixNano())

	max, _ := strconv.Atoi(os.Args[1])

	// declare an uninitialized nil slice
	var uniques []int

loop:
	// you can still use the len function on a nil slice
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		// append function can add new elements to a slice
		uniques = append(uniques, n)

		// a slice doesn't contain any elements right from the start
		// uniques[found] = n
		// found++
	}

	fmt.Println("\n\nuniques:", uniques)
	fmt.Println("\nlength of uniques:", len(uniques))

	sort.Ints(uniques)
	fmt.Println("\nsorted:", uniques)

	nums := [5]int{5, 4, 3, 2, 1}
	sort.Ints(nums[:])
	fmt.Println("\nnums:", nums)
}
