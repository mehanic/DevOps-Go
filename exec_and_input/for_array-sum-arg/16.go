package main

import "fmt"

func main() {
	var n int
	fmt.Print("n=")
	fmt.Scan(&n)
	var numbers [100]int
	for i := 0; i < n; i++ {
		var number int
		fmt.Print("number=")
		fmt.Scan(&number)
		numbers[i] = number
	}
	sumi := 0
	for i := 0; i < n; i++ {
		sumi = sumi + numbers[i]
	}
	fmt.Println(sumi)
}

// n=3
// number=56
// number=78
// number=36
// 170
