package main

import "fmt"

// Car struct represents a basic car
type Car struct {
	Make  string
	Model string
	Year  int
}

// ElectricCar struct embeds Car and has an additional Battery
type ElectricCar struct {
	Car
	Battery Battery
}

// Battery struct represents the battery of an electric car
type Battery struct {
	BatterySize int
}

// Method to get the descriptive name of the car
func (c Car) GetDescriptiveName() string {
	return fmt.Sprintf("%d %s %s", c.Year, c.Make, c.Model)
}

// Method to describe the battery
func (b Battery) DescribeBattery() {
	fmt.Printf("This car has a %d-kWh battery.\n", b.BatterySize)
}

// Method to get the range of the electric car based on battery size
func (b Battery) GetRange() {
	var rangeMiles int
	if b.BatterySize == 75 {
		rangeMiles = 260
	} else if b.BatterySize == 100 {
		rangeMiles = 315
	}
	fmt.Printf("This car can go approximately %d miles on a full charge.\n", rangeMiles)
}

func main() {
	// Create an instance of ElectricCar with a 100 kWh battery
	myTesla := ElectricCar{
		Car: Car{
			Make:  "Tesla",
			Model: "Model S",
			Year:  2016,
		},
		Battery: Battery{
			BatterySize: 100,
		},
	}

	// Get the descriptive name of the electric car
	fmt.Println(myTesla.GetDescriptiveName())

	// Describe the battery and get the range
	myTesla.Battery.DescribeBattery()
	myTesla.Battery.GetRange()
}

// 2016 Tesla Model S
// This car has a 100-kWh battery.
// This car can go approximately 315 miles on a full charge.
