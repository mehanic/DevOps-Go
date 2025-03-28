package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set a[4]=100:", a)
	fmt.Println("a[4]:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

// emp: [0 0 0 0 0]
// set a[4]=100: [0 0 0 0 100]
// a[4]: 100
// len: 5
// dcl: [1 2 3 4 5]
// 2d:  [[0 1 2] [1 2 3]]
