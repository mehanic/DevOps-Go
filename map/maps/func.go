package maps

import "fmt"

func CreateMap() {

	/*
		Maps are used to store data values in key:value pairs.

		Each element in a map is a key:value pair.

		A map is an unordered and changeable collection that does not allow duplicates.

	*/

	capitals := make(map[string]string) // We create a map by giving keys and values

	// We can also create like that :
	//capitals:= map[string]string{"Turkey": "Ankara", "Germany": "Berlin", "France": "Paris", "Austria": "Vienna"}

	capitals["Turkey"] = "Ankara"
	capitals["Germany"] = "Berlin"
	capitals["France"] = "Paris"
	capitals["Austria"] = "Vienna"

	fmt.Println(capitals)
	fmt.Println("Number of Elements : ", len(capitals))

	// We can delete an Element with Delete Function

	delete(capitals, "France")
	fmt.Println(capitals)
	fmt.Println("Number of Elements : ", len(capitals))

	// We can check if a key exist in a map

	value, is_exist := capitals["Turkey"]
	fmt.Println("Is value", value, "is in the list? : ", is_exist)

	ne_value, ne_is_exist := capitals["Netherlands"]
	fmt.Println("Is value", ne_value, "is in the list? : ", ne_is_exist)

}
