package main

import "fmt"

func main() {

	// Switch

	grade := -7
	switch grade { // Similar to if-else, can also be written as: switch grade := 3 {}
	default:
		fmt.Println("Invalid Grade")
	case 5: // if grade == 5 { fmt.Println("Excellent") }
		fmt.Println("Excellent")
	case 4:
		fmt.Println("Good")
	case 3:
		fmt.Println("Average")
	case 2:
		fmt.Println("Pass")
		y := 100
		fmt.Println(y)
	case 1:
		fmt.Println("Fail")
	}

	// If no case is met, it will go to the default
	// Data type is important as well.
	// break is not manually written; it behaves as if it exists.

	switch {
	case false:
		fmt.Println("This will not appear in the console.")
	case true:
		fmt.Println("This will appear in the console.")
	case true:
		fmt.Println("This will not appear either") // because there is an imaginary break here
	}

	// LOOPS
	// Only 'for' loop exists in Go; there is no 'while', but 'for' can be used to create a while loop.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	// For is Go's "while"
	fmt.Println(sum)

	/* 	for i := 0; i <= 10; i += 5 {
		fmt.Println(i)
	} */

	// for <init statement>; <condition>; <post statement> { ----- }

	/* 	i := 0
	   	for ; i <= 10; i += 5 {
	   		fmt.Println(i)
	   	}
	   	fmt.Println(i) */

	/* 	for {  // Infinite Loop
		fmt.Println("My name is Arin")
	} */

	/* 	for i := 0; true; i += 5 {
		fmt.Println(i)
	} */

	/* 	for i := 0; false; i += 5 {
		fmt.Println(i)
	} */

	/* 	i := 10
	   	for i >= 0 {
	   		fmt.Println(i)
	   		i--
	   	} */

	/* 	for i := 0; i <= 10; i++ { // continue --> go to the beginning of the loop
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	} */

	for i := 0; i <= 10; i++ {

		if i == 3 {
			break
		}

		fmt.Println(i)
	} // break --> exit the loop
	// continue skips that step

	// APPLICATION

	// fallthrough keyword --> skips the break keyword

	/*switch x := 75; {
	case x < 20:
		fmt.Printf("%d is smaller than 20\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d is smaller than 50\n", x)
		fallthrough
	case x < 100:
		fmt.Printf("%d is smaller than 100\n", x)
		fallthrough
	case x < 200:
		fmt.Printf("%d is smaller than 200\n", x)
	}*/

	// Both 100 and 200 will be printed.

	// Writing idiomatic code in Go means writing parts of code in an independent, simple, and readable manner.

	// Less idiomatic
	if x := 20; x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}

	// More idiomatic
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "is even")
		return
	}
	fmt.Println(x, "is odd")

}

// 5-) Write a program that shows prime numbers between 1 and 50.

func prime() {

	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf("%d is a prime number\n", x)
		}
	}
}


// Invalid Grade
// This will appear in the console.
// 1024
// 0
// 1
// 2
// 20 is even
// 20 is even
