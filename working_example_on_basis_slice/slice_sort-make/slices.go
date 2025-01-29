package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Working with slices")

	//initializing a slice
	var arr = []string{}
	fmt.Printf("Type of slice is %T\n", arr)

	arr = append(arr, "cricker", "football", "tennis")
	fmt.Println(arr)

	newarr := arr[1:]
	fmt.Println(newarr)

	highScores := make([]int, 5)
	highScores[0] = 45
	highScores[1] = 18
	highScores[2] = 7
	highScores[3] = 228
	highScores[4] = 333
	fmt.Println(highScores)
	highScores = append(highScores, 27)
	sort.Ints(highScores)
	fmt.Println(highScores)
	var a int = 1
	highScores = append(highScores[:a], highScores[a+1:]...)
	fmt.Println(highScores)

}
