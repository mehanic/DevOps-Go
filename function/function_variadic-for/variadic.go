package main

import "fmt"

// Variadic function
func promedio(marks ...int) int {
	var promedio, suma int
	for _, mark := range marks {
		suma += mark
	}
	promedio = suma / len(marks)
	return promedio
}

func main() {
	// Declare result variable
	var result int

	// Calling the variadic function with a slice
	nums := []int{1, 2, 3, 4}
	result = promedio(nums...)
	fmt.Println("Variadic result:", result)
}

// Variadic result: 2
