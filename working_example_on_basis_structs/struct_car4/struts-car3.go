package main

import (
	"fmt"
)

// Car represents a simple car with attributes and methods.
type Car struct {
	make             string
	model            string
	year             int
	odometerReading  int
}

// NewCar creates a new Car instance.
func NewCar(make, model string, year int) *Car {
	return &Car{make: make, model: model, year: year, odometerReading: 20}
}

// GetDescriptiveName returns a neatly formatted descriptive name.
func (c *Car) GetDescriptiveName() string {
	return fmt.Sprintf("%d %s %s", c.year, c.make, c.model)
}

// ReadOdometer prints a statement showing the car's mileage.
func (c *Car) ReadOdometer() {
	fmt.Printf("This car has %d miles on it.\n", c.odometerReading)
}

// UpdateOdometer sets the odometer reading to the given value.
func (c *Car) UpdateOdometer(mileage int) {
	if mileage >= c.odometerReading {
		c.odometerReading = mileage
	} else {
		fmt.Println("You can't roll back an odometer!")
	}
}

// IncrementOdometer adds the given amount to the odometer reading.
func (c *Car) IncrementOdometer(miles int) {
	c.odometerReading += miles
}

// Battery represents a battery for an electric car.
type Battery struct {
	batterySize int
}

// NewBattery creates a new Battery instance with a default size.
func NewBattery(batterySize int) *Battery {
	return &Battery{batterySize: batterySize}
}

// DescribeBattery prints a statement describing the battery size.
func (b *Battery) DescribeBattery() {
	fmt.Printf("This car has a %d-kwh battery.\n", b.batterySize)
}

// ElectricCar represents an electric car, inheriting from Car.
type ElectricCar struct {
	*Car
	battery *Battery
}

// NewElectricCar creates a new ElectricCar instance.
func NewElectricCar(make, model string, year int) *ElectricCar {
	return &ElectricCar{
		Car:     NewCar(make, model, year),
		battery: NewBattery(70),
	}
}

// DescribeBattery prints a statement describing the battery size.
func (e *ElectricCar) DescribeBattery() {
	e.battery.DescribeBattery()
}

// FillGasTank prints a message indicating that electric cars don't need gas tanks.
func (e *ElectricCar) FillGasTank() {
	fmt.Println("This car doesn't need a gas tank!")
}

func main() {
	myTesla := NewElectricCar("Tesla", "Model S", 2016)
	fmt.Println(myTesla.GetDescriptiveName())
	myTesla.DescribeBattery()
}

// 2016 Tesla Model S
// This car has a 70-kwh battery.
