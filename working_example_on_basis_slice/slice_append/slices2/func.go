package slices

import (
	"fmt"
	"math"
)

func CreateSlice() {

	/*
		Slices are similar to arrays, but are more powerful and flexible.

		Like arrays, slices are also used to store multiple values of the same type in a single variable.

		However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	*/

	fmt.Println("How many cities do you want to go?")

	var city_number int

	fmt.Scanln(&city_number)

	cities := make([]string, 0)

	// or cities:= []string{}

	for ind := 0; ind < city_number; ind++ {

		fmt.Println("The city you want to visit in the order", ind+1)

		var city string

		fmt.Scanln(&city)

		cities = append(cities, city) // we use append function to append elements to slice.

	}

	fmt.Println("------------------------------")
	fmt.Println("\nLength of the slice: ", len(cities))   // returns the length of the slice (the number of elements in the slice)
	fmt.Println("\nCapacity of the slice: ", cap(cities)) // returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	fmt.Println("------------------------------")
	fmt.Println("\nCities you want to see \n", cities)
	fmt.Println("\nCities you want to see most \n", cities[0:3])
	fmt.Println("\nCities you want to see least \n", cities[city_number-3:])
}

func CreateSliceFromArray() {

	fmt.Println("How many numbers do you want to see?")

	var number_selection float64

	fmt.Scanln(&number_selection)

	numbers := []float64{}

	for ind := 0.0; ind < number_selection; ind++ {

		numbers = append(numbers, math.Pow(ind, 2)) // we use append function to append elements to slice.

	}

	fmt.Println("------------------------------")

	fmt.Println("This is the array you've created", numbers)
	fmt.Println("------------------------------")
	fmt.Println("\nSelect you slice beginning value")
	var slice_beginning int
	fmt.Scanln(&slice_beginning)

	fmt.Println("Select you slice ending value")
	var slice_ending int
	fmt.Scanln(&slice_ending)

	fmt.Println("Here is our new slice", numbers[slice_beginning:slice_ending])

}
