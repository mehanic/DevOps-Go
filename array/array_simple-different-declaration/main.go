package main

import "fmt"

func main() {
	cityOne := "istanbul"
	cityTwo := "roma"
	cityThree := "tahran"
	cityFour := "belgrad"
	fmt.Println(cityOne, cityTwo, cityThree, cityFour)

	town := [4]string{"Tokio", "Seul", "Sapporo", "Kioto"}
	fmt.Println(town)

	cities1 := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	for i := 0; i < len(cities1); i++ {
		fmt.Println(i, cities1[i])
	}
	cities1[0] = "ANKARA"
	fmt.Println()
	for i := 0; i < len(cities1); i++ {
		fmt.Println(i, cities1[i])
	}

	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	for _, city := range cities {
		fmt.Println(city)
	}

	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myArr = mySquare(myArr) // First Class Functions
	fmt.Println(myArr)
}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
