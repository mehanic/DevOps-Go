package main

import "fmt"

func Array3() {

	sayilar := [5]int{233, 24, 56, 87, 23}
	fmt.Println(sayilar)
	fmt.Println("-------------------")
	enBüyük := sayilar[0] //The variable enBüyük is initialized to the first element of the array, sayilar[0], which is 233. This variable will store the maximum value found during iteration.
	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
		if sayilar[i] > enBüyük {
			enBüyük = sayilar[i]
		}

	}
	fmt.Println(enBüyük)
	fmt.Println("-------------------")
}

func Array4() {
	var sayilar [2][3]int
	sayilar[0][0] = 12
	sayilar[0][1] = 31
	sayilar[0][2] = 98
	sayilar[1][0] = 23
	sayilar[1][1] = 54
	sayilar[1][2] = 13
	fmt.Println(sayilar)
}

func main() {
	Array3()
	Array4()
}

// [233 24 56 87 23]
// -------------------
// 233
// 24
// 56
// 87
// 23
// 233
// -------------------
// [[12 31 98] [23 54 13]]
