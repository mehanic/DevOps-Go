package main

import "fmt"

// Exercise 1
func Exercise1() {

	// - Create a "person" type with a struct underlying type that can contain the following fields:
	// - Name
	// - Surname
	// - Favorite ice cream flavors
	// - Create two values of the "person" type and demonstrate these values using range over the slice that contains the ice cream flavors.

	type Person struct {
		name      string
		surname   string
		favorites []string
	}

	person1 := Person{
		name:    "Micael",
		surname: "Ramos",
		favorites: []string{"Chocolate", "Vanilla", "Strawberry"},
	}

	fmt.Println(person1)
}

// Exercise 2
func Exercise2() {

	// - Using the previous solution, place the "person" type values in a map, using surnames as keys.
	// - Demonstrate the values of the map using range.
	// - The different flavors should be demonstrated using another range inside the previous range.

	type Person struct {
		name      string
		surname   string
		favorites []string
	}

	users := make(map[string]Person)

	users["Ramos"] = Person{
		name:    "Micael",
		surname: "Ramos",
		favorites: []string{"Chocolate", "Strawberry", "Pixi"},
	}

	users["Santana"] = Person{
		"Isabelli", 
		"Santana", 
		[]string{"Micael", "Micael", "Micael"}}

	for key, value := range users {
		fmt.Println(key)
		for _, flavor := range value.favorites {
			fmt.Print(flavor, " ")
		}
		fmt.Println(" ")
	}
}

// Exercise 3
func Exercise3() {

	// - Create a new type: vehicle
	// - The underlying type should be struct
	// - It should contain the fields: doors, color
	// - Create two new types: pickup and sedan
	// - The underlying types should be struct
	// - Both should contain "vehicle" as an embedded struct
	// - The pickup type should contain a bool field called "fourWheelDrive"
	// - The sedan type should contain a bool field called "luxuryModel"
	// - Using the structs vehicle, pickup, and sedan:
	// - Using composite literal, create a value of type pickup and assign values to its fields
	// - Using composite literal, create a value of type sedan and assign values to its fields
	// - Demonstrate these values.
	// - Demonstrate one field from each of the two types.

	type Vehicle struct {
		doors int
		color string
	}

	type Pickup struct {
		Vehicle
		fourWheelDrive bool
		luxuryModel    bool
	}

	type Sedan struct {
		Vehicle
		fourWheelDrive bool
		luxuryModel    bool
	}

	fiatStrada := Pickup{
		Vehicle: Vehicle{
			doors: 4,
			color: "Red",
		},
		fourWheelDrive: false,
		luxuryModel:    false,
	}

	toyotaCorolla := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "Gray",
		},
		fourWheelDrive: false,
		luxuryModel:    false,
	}

	fmt.Println(fiatStrada)
	fmt.Println(toyotaCorolla)

	// Demonstrating specific fields
	fmt.Println(fiatStrada.Vehicle)
	fmt.Println(toyotaCorolla.Vehicle)
}

// Exercise 4
func Exercise4() {

	// - Create and use an anonymous struct.
	// - Challenge: Inside the struct, have a value of type map and another of type slice.

	Customer := struct {
		name       string
		vehicles   map[string]string
		family     []string
	}{
		name: "Fernando",
		vehicles: map[string]string{
			"uncle's car":   "sedan",
			"nephew's car":  "beetle",
		},
		family: []string{"wife", "son", "mother"},
	}

	fmt.Print(Customer)
}

func main() {
	// Call the exercises here to test
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
}

// {Micael Ramos [Chocolate Vanilla Strawberry]}
// Ramos
// Chocolate Strawberry Pixi  
// Santana
// Micael Micael Micael  
// {{4 Red} false false}
// {{4 Gray} false false}
// {4 Red}
// {4 Gray}
// {Fernando map[nephew's car:beetle uncle's car:sedan] [wife son mother]}
