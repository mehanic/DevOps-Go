package main

import "fmt"

func main() {

	// simple
	fmt.Println("simple for-loop")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// nested
	fmt.Println("nested-loop")
	for i := 0; i < 2; i++ {
		//outer loop
		fmt.Println("outer loop:", i)
		for j := 0; j < 3; j++ {
			// inner loop
			fmt.Println("\tinner loop:", j)
		}
	}

	// like C while, no keyword while in Go
	fmt.Println("like C while")
	var x = 0
	for x < 10 {
		fmt.Print(x, " ")
		x++
	}
	fmt.Println()

	// for with break
	fmt.Println("for with break")
	x = 0
	for x < 10 {
		if x == 5 {
			break // break will kick you out of loop
		}
		fmt.Print(x, " ")
		x++
	}
	fmt.Println()

	// for with continue
	fmt.Println("for with continue")
	x = 0
	for x < 10 {
		x++
		if x%2 == 0 {
			continue // continue 會跳過這次回圈剩餘未執行的程式碼，直接進行下一個迴圈。
		}
		fmt.Print(x, " ")
	}
	fmt.Println()
}
