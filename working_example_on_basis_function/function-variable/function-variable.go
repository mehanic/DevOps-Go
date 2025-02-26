package main

import "fmt"

func main() {
	// Declaring two int values num1 and num2 with values 10, 5 resepectively
	num1, num2 := 10, 5

	// Assigning function sub to a varible f1.
	f1 := sub
	fmt.Printf("\nType of f1 : %T\t\tResult: %d", f1, f1(num1, num2))
	//Output : Type of f1 : func(int, int) int		Result: 5

	// Assigning function add a varible f2.
	f2 := add
	fmt.Printf("\nType of f2 : %T\t\tResult: %d", f2, f2(num1, num2))
	//Output : Type of f2 : func(int, int) int		Result: 15

	// Assigning function mul a varible f3.
	f3 := mul
	fmt.Printf("\nType of f3 : %T\t\tResult: %d", f3, f3(num1, num2))
	//Output : Type of f3 : func(int, int) int		Result: 50

	// Assigning function div a varible f4.
	f4 := div
	fmt.Printf("\nType of f4 : %T\t\tResult: %d", f4, f4(num1, num2))
	//Output : Type of f4 : func(int, int) int		Result: 2

}
func sub(num1, num2 int) int {
	return num1 - num2
}
func add(num1, num2 int) int {
	return num1 + num2
}
func mul(num1, num2 int) int {
	return num1 * num2
}
func div(num1, num2 int) int {
	return num1 / num2
}
