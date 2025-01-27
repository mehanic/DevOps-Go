package main

import "fmt"

/*
control statements
if /else
if/elseif /else
switch statement
*/

func main() {

	a := 4
	b := 5

	// Comparison
	fmt.Printf("a == b, %v\n", a == b) // false
	fmt.Printf("a == a, %v\n", a == a) // true
	fmt.Printf("a != b, %v\n", a != b) // true
	fmt.Printf("a > b, %v\n", a > b)   // false
	fmt.Printf("a < b, %v\n", a < b)   // true
	fmt.Printf("a <= b, %v\n", a <= b) // true
	fmt.Printf("a <= b, %v\n", a <= a) // true
	fmt.Printf("a >= b, %v\n", a >= a) // true

	fmt.Println("--------------------")

	// if without else statement
	if a > b {
		fmt.Printf("if statement: a is greater than b , %v\n", a > b)
	}
	fmt.Printf("else statement: a is lesser than b , %v\n", a > b)

	// if - else if - else statement
	if a > b {
		fmt.Printf("if: a is greater than b , %v\n", a > b)
	} else if a == b {
		fmt.Printf("if else: a is same as  b , %v\n", a > b)
	} else {
		fmt.Printf("else: a is lesser than b , %v\n", a > b)
	}
	fmt.Println("--------------------")
	//switch case
	switch {
	case a < 5:
		fmt.Println("switch: a is less than 5")
	case a == 5:
		fmt.Println("switch: a is equal to 5")
	default:
		fmt.Println("switch: a is greater than 5")
	}

	switch a {
	case 5:
		fmt.Println("switch: a is equal to 5")
	}

	// for loop
	for a < 10 {
		fmt.Printf("For: a - %v\n", a)
		a++
	}

	for i := 0; i < b; i++ {
		fmt.Printf("for loop b: %v\n", i)
	}

}

// $ go run /home/mehanic/structure/if-else/if_else-switch-for/control.go
// a == b, false
// a == a, true
// a != b, true
// a > b, false
// a < b, true
// a <= b, true
// a <= b, true
// a >= b, true
// else statement: a is lesser than b , false
// else: a is lesser than b , false
// switch: a is less than 5
// For: a - 4
// For: a - 5
// For: a - 6
// For: a - 7
// For: a - 8
// For: a - 9
// for loop b: 0
// for loop b: 1
// for loop b: 2
// for loop b: 3
// for loop b: 4
