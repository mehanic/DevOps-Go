package main

import (
	"fmt"
	"strings"
)

// Restaurant struct definition
type Restaurant struct {
	restaurantName string
	cuisineType     string
}

// Method to describe the restaurant
func (r Restaurant) describeRestaurant() {
	fmt.Println("This restaurant is called " + strings.ToUpper(r.restaurantName) + ".")
	fmt.Println("The type of this restaurant is " + strings.Title(r.cuisineType) + ".")
}

// Method to open the restaurant
func (r Restaurant) openRestaurant() {
	fmt.Println("Our restaurant is opening.")
}

func main() {
	// Create an instance of Restaurant
	KFC := Restaurant{restaurantName: "kfc", cuisineType: "fast food"}

	// Call methods on the instance
	KFC.describeRestaurant()
	KFC.openRestaurant()
}

// This restaurant is called KFC.
// The type of this restaurant is Fast Food.
// Our restaurant is opening.

