package main

import "fmt"

// Function to display a value with a default argument
func display(a ...int) {
	value := 1 // default value
	if len(a) > 0 {
		value = a[0]
	}
	fmt.Println("The value of a is:", value)
}

func main() {
	display(4)
	display(5)
	display()
	main2()
	main3()
}

// Function to add two numbers with a default value for the second argument
func addNumbers(a int, b ...int) {
	var result int
	if len(b) > 0 {
		result = a + b[0]
	} else {
		result = a
	}
	fmt.Println("The result is:", result)
}

func main2() {
	addNumbers(4, 5)
	addNumbers(5)
	addNumbers(7)
}

// Function to work with a user with a default value
func workingOnSome(user ...string) {
	username := "root" // default value
	if len(user) > 0 {
		username = user[0]
	}
	fmt.Printf("working with %s\n", username)
}

func main3() {
	workingOnSome("weblogic_admin")
}

// go run /home/mehanic/structure/function/function17/function17.go
// The value of a is: 4
// The value of a is: 5
// The value of a is: 1
// The result is: 9
// The result is: 5
// The result is: 7
// working with weblogic_admin
