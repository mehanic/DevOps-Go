package main

import (
	"fmt"
	"math"
)

func main() {
	var length, width, radius float64

	// Prompt user for input
	fmt.Print("Please enter the length, width, and radius: ")
	_, err := fmt.Scanf("%f %f %f", &length, &width, &radius) // Read multiple float inputs
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Calculate area and perimeter of rectangle
	areaRectangle := length * width
	perimeterRect := 2 * (length + width)

	// Calculate area and circumference of circle
	areaCircle := math.Pi * radius * radius
	circumferenceCircle := 2 * math.Pi * radius

	// Print results
	fmt.Printf("Area of rectangle = %.2f\n", areaRectangle)
	fmt.Printf("Perimeter of Rectangle = %.2f\n", perimeterRect)
	fmt.Printf("Area of circle = %.2f\n", areaCircle)
	fmt.Printf("Circumference of circle = %.2f\n", circumferenceCircle)
}

