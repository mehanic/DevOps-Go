package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const (
		total   = 500
		inGroup = 20
		maxGroups = 25
	)

	people := preparePeople(total)

	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})

	groups := toGroups(people, inGroup)

	fmt.Println(groups)
	for i := 0; i < maxGroups-1; i++ {
		groups = shift(groups, inGroup)

		fmt.Println(groups)
	}
}

func shift(groups [][]int, inGroup int) [][]int {
	result := make([][]int, len(groups))

	for i := 0; i < len(groups); i++ {
		item := make([]int, inGroup)

		for j := 0; j < inGroup; j++ {
			y := (j + i) % len(groups)

			item[j] = groups[y][j]

		}

		result[i] = item
	}

	return result
}

func preparePeople(total int) []int {
	people := make([]int, total)
	for i := 0; i < len(people); i++ {
		people[i] = i + 1
	}

	return people
}

func toGroups(people []int, inGroup int) [][]int {
	groups := make([][]int, len(people)/inGroup)
	for i := 0; i < len(groups); i++ {
		groups[i] = people[i*inGroup : (i+1)*inGroup]
	}

	return groups
}


