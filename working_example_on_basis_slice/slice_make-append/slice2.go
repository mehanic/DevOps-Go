package main

import "fmt"

func main() {
	// Initial slices
	myFoods := []string{"pizza", "falafel", "carrot cake"}
	friendFoods := make([]string, len(myFoods))
	copy(friendFoods, myFoods)

	// Modify slices
	myFoods = append(myFoods, "cannoli")
	friendFoods = append(friendFoods, "ice cream")

	// Print the results
	fmt.Println("My favorite foods are:")
	for _, food := range myFoods {
		fmt.Println(food)
	}

	fmt.Println("\nMy friend's favorite foods are:")
	for _, food := range friendFoods {
		fmt.Println(food)
	}
}

