package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}
	sumi := 0
	for i := 0; i < 5; i++ {
		sumi = sumi + numbers[i]
	}
	fmt.Println(sumi)
}

// 15
