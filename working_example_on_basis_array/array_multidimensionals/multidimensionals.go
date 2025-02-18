package arrays

import "fmt"

func CreateMultidimensionalArray() {

	var countries_and_cities [3][2]string

	countries_and_cities[0][0] = "Istanbul"
	countries_and_cities[0][1] = "Ankara"
	countries_and_cities[1][0] = "Berlin"
	countries_and_cities[1][1] = "Hamburg"
	countries_and_cities[2][0] = "Vienna"
	countries_and_cities[2][1] = "Graz"

	fmt.Println("Countries and Cities ", countries_and_cities)

	for ind := 0; ind < 3; ind++ {
		for sind := 0; sind < 2; sind++ {
			fmt.Println(countries_and_cities[ind][sind])
		}
	}

}
