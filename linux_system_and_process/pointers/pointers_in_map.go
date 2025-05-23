package pointers

import "fmt"

func PointersInMap() {

	capitals := map[string]string{"Turkey": "Ankara", "Austria": "Vienna"}
	b := capitals

	fmt.Println(capitals, b)

	capitals["Turkey"] = "Istanbul"
	fmt.Println(capitals, b)

	fmt.Println("It contains internal pointrs, so copies point to same underlying data.")

}
