package main

import (
	"fmt"
	"sort"
	"strings"
)

// Custom type for reverse sorting
type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] > s[j] } // Reverse order
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// Initial slice
	cars := []string{"bmw", "audi", "toyota", "subaru"}

	// Sort in reverse order
	sort.Sort(StringSlice(cars))
	fmt.Println(cars)

	// Print original list (after sorting)
	fmt.Println("Here is the original list:")
	fmt.Println(cars)

	// Print sorted list
	sortedCars := make([]string, len(cars))
	copy(sortedCars, cars)
	sort.Strings(sortedCars)
	fmt.Println("\nHere is the sorted list:")
	fmt.Println(sortedCars)

	// Print original list again
	fmt.Println("\nHere is the original list again:")
	fmt.Println(cars)

	// Reverse the list
	for i := 0; i < len(cars)/2; i++ {
		cars[i], cars[len(cars)-1-i] = cars[len(cars)-1-i], cars[i]
	}
	fmt.Println(cars)
	fmt.Println(len(cars))

	// Print last element of motorcycles slice
	motorcycles := []string{"honda", "yamaha", "suzuki"}
	fmt.Println(motorcycles[len(motorcycles)-1])

	// Iterate through cars slice and print formatted strings
	cars = []string{"audi", "bmw", "subaru", "toyota"}
	for _, car := range cars {
		if car == "bmw" {
			fmt.Println(strings.ToUpper(car))
		} else {
			fmt.Println(strings.Title(car))
		}
	}
}

// [toyota subaru bmw audi]
// Here is the original list:
// [toyota subaru bmw audi]

// Here is the sorted list:
// [audi bmw subaru toyota]

// Here is the original list again:
// [toyota subaru bmw audi]
// [audi bmw subaru toyota]
// 4
// suzuki
// Audi
// BMW
// Subaru
// Toyota
