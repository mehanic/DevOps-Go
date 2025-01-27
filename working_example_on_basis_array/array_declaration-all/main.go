package main

import "fmt"

func main() {
	var numbersZero [4]int

	fmt.Printf("%d\n", numbersZero)
	fmt.Printf("%#v\n", numbersZero)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [4]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	//ellipsis operator

	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The length of a5 is %d\n", len(a5))

	for i, v := range a2 {
		fmt.Println("Index: ", i, "Values: ", v)
	}

	//multi-dimensional arrays

	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10}, //this comma is mandatory
	}
	fmt.Println(balances)

	// keyed elements array
	grades := [3]int{
		1: 10,
		0: 3,
		2: 43,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{5: "Willkommen"}
	fmt.Printf("%#v\n", names)

	cities := [...]string{
		5: "Kaduna",
		"London",
		1: "France",
	}
	fmt.Printf("%#v\n", cities)

	numbers := []int{2, 3, 4, 5}

	fmt.Println(numbers)
}

// [0 0 0 0]
// [4]int{0, 0, 0, 0}
// [4]float64{0, 0, 0, 0}
// [4]int{-10, 1, 100, 0}
// [4]string{"Dan", "Diana", "Paul", "John"}
// [4]string{"x", "y", "", ""}
// [6]int{1, 2, 5, 1, -10, 66}
// The length of a5 is 6
// Index:  0 Values:  -10
// Index:  1 Values:  1
// Index:  2 Values:  100
// Index:  3 Values:  0
// [[5 6 7] [8 9 10]]
// [3 10 43]
// [0 0 50]
// [6]string{"", "", "", "", "", "Willkommen"}
// [7]string{"", "France", "", "", "", "Kaduna", "London"}
// [2 3 4 5]
