package main

import "fmt"

func main() {
	rates := [3]float64{
		// index 0 empty
		// index 1 empty
		2: 1.5, // index: 2
	}

	fmt.Println(rates)
	main1()
	main2()
}

//--

func main1() {
	// below ellipsis (...) calculates the length of the
	// array automatically
	rates := [...]float64{
		// index 0 empty
		// index 1 empty
		// index 2 empty
		// index 3 empty
		// index 4 empty
		5: 1.5, // index: 5
	}

	fmt.Println(rates)

}

//---

func main2() {
	rates := [...]float64{
		// index 1 to 4 empty

		5:   1.5, // index: 5
		2.5, //      index: 6
		0:   0.5, // index: 0
	}

	fmt.Println(rates)
}

// [0 0 1.5]
// [0 0 0 0 0 1.5]
// [0.5 0 0 0 0 1.5 2.5]
