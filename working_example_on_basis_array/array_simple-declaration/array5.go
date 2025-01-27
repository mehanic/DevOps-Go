package main

import "fmt"

func main() {
	rates := [3]float64{
		0.5,
		2.5,
		1.5,
	}

	fmt.Println(rates)
	main2()

}

func main2() {
	rates := [3]float64{
		0: 0.6, // index: 0
		1: 2.6, // index: 1
		2: 1.6, // index: 2
	}

	fmt.Println(rates)
}

// [0.5 2.5 1.5]
// [0.6 2.6 1.6]
