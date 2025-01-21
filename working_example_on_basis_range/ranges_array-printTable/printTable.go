package main

import (
	"fmt"
)

func main() {
	tableData := [][]string{
		{"apples", "oranges", "cherries", "bananas"},
		{"Alice", "Bob", "Carol", "David"},
		{"dogs", "cats", "moose", "goose"},
	}

	printTable(tableData)
}

func printTable(dataLists [][]string) {
	colWidths := make([]int, len(dataLists))
	for i := range colWidths {
		for _, data := range dataLists {
			if len(data[i]) > colWidths[i] {
				colWidths[i] = len(data[i])
			}
		}
	}

	for x := range dataLists[0] {
		for i, data := range dataLists {
			fmt.Print(data[x])
			if i < len(dataLists)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// apples Alice dogs
// oranges Bob cats
// cherries Carol moose
// bananas David goose
