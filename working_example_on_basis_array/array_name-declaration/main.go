package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Veges list is: ", vegList)
	fmt.Println("Veges list length is: ", len(vegList))
}

// Welcome to a class on arrays
// Fruit list is:  [Apple Tomato  Peach]
// Fruit list length is:  4
// Veges list is:  [potato beans mushroom]
// Veges list length is:  3
