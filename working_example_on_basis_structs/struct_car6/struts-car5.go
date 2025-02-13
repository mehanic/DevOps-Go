package main

import (
	"fmt"
	"strconv"
)

// Car struct represents a regular car
type Car struct {
	make  string
	model string
	year  int
}

// ElectricCar struct embeds the Car struct, representing an electric car
type ElectricCar struct {
	Car          // Embedded Car struct to reuse its fields and methods
	batteryRange int
}

// Method to get the descriptive name of the car (regular or electric)
func (c *Car) GetDescriptiveName() string {
	return strconv.Itoa(c.year) + " " + c.make + " " + c.model
}

// Main function to demonstrate Car and ElectricCar usage
func main() {
	// Creating a regular car instance
	myBeetle := Car{
		make:  "volkswagen",
		model: "beetle",
		year:  2016,
	}
	fmt.Println(myBeetle.GetDescriptiveName())

	// Creating an electric car instance
	myTesla := ElectricCar{
		Car: Car{
			make:  "tesla",
			model: "roadster",
			year:  2016,
		},
		batteryRange: 400, // For demonstration purposes, can add more attributes for electric cars
	}
	fmt.Println(myTesla.GetDescriptiveName())
}

// 2016 volkswagen beetle
// 2016 tesla roadster
