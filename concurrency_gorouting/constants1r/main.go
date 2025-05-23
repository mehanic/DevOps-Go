package main

import "fmt"

func main() {
	// const a int = 3
	// _ = a

	const (
		min1 = 500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)

	const (
		c1 = iota
		c2
		c3
	)

	fmt.Println(c1, c2, c3)

	const (
		x = -(iota + 2)
		_ // -3 will be skipped
		y
		z
	)

	fmt.Println(x, y, z)

	// Understanding arrays

	var numbers = [4]int{4, 5, 6, 7}
	fmt.Println(numbers)

	var cities = [4]string{"New York", "London", "Paris", "Singapore"}
	fmt.Printf("%T\n", cities)

	balances := map[string]float64{
		"USD": 233.2,
		"EUR": 555.2,
	}

	fmt.Printf("%T\n", balances)
	fmt.Println(balances)

	// STRUCT TYPE
	type Person struct {
		name string
		age  int
	}

	var you Person

	fmt.Printf("%T\n", you)

	//POINTER TYPE
	var p int = 2
	ptr := &p
	fmt.Printf("ptr is of type %T with value %v and address %v\n", ptr, ptr, &ptr)
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 13:56:20
// 500 500 500
// 0 1 2
// -2 -4 -5
// [4 5 6 7]
// [4]string
// map[string]float64
// map[EUR:555.2 USD:233.2]
// main.Person
// ptr is of type *int with value 0xc00011a068 and address 0xc00011e040
