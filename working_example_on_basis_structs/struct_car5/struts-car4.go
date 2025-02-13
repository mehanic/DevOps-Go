package main

import (
	"fmt"
	"strconv"
)

// Car struct
type Car struct {
	make            string
	model           string
	year            int
	odometerReading int
}

// Method to get the descriptive name of the car
func (c *Car) GetDescriptiveName() string {
	return strconv.Itoa(c.year) + " " + c.make + " " + c.model
}

// Method to read the odometer reading
func (c *Car) ReadOdometer() {
	fmt.Println("This car has", c.odometerReading, "miles on it.")
}

// Main function to demonstrate the usage of the Car struct
func main() {
	// Creating a new car object
	myNewCar := Car{
		make:  "audi",
		model: "a4",
		year:  2016,
	}

	// Print descriptive name
	fmt.Println(myNewCar.GetDescriptiveName())

	// Set odometer reading
	myNewCar.odometerReading = 23

	// Print odometer reading
	myNewCar.ReadOdometer()
}

// 2016 audi a4
// This car has 23 miles on it.
