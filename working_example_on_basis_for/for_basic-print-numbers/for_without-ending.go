package main

import "fmt"

func main() {
	//sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	sum := 0
	for sum < 1000 {
		sum += 100
	}
	fmt.Println(sum)
	k := 0
	for k < 100 {
		k += 10
		fmt.Println(k)
	}
	//without ending
	// for true {
	// 	fmt.Println(11)
	// }

	// for {
	// 	fmt.Println(11)
	// }
}

// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 1000
// 10
// 20
// 30
// 40
// 50
// 60
// 70
// 80
// 90
// 100
