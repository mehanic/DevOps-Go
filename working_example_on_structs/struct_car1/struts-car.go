package main

import (
    "fmt"
    "strings"
)

type Car struct {
    Make              string
    Model             string
    Year              int
    OdometerReading   int
}

func (c *Car) GetDescriptiveName() string {
    longName := fmt.Sprintf("%d %s %s", c.Year, c.Make, c.Model)
    return strings.Title(longName)
}

func (c *Car) ReadOdometer() {
    fmt.Printf("This car has %d miles on it.\n", c.OdometerReading)
}

func (c *Car) UpdateOdometer(mileage int) {
    if mileage >= c.OdometerReading {
        c.OdometerReading = mileage
    } else {
        fmt.Println("You can't roll back an odometer!")
    }
}

func (c *Car) IncrementOdometer(miles int) {
    c.OdometerReading += miles
}

//---
type Battery struct {
    BatterySize int
}

func (b *Battery) DescribeBattery() {
    fmt.Printf("This car has a %d-kwh battery.\n", b.BatterySize)
}

func (b *Battery) GetRange() {
    var rangeMiles int
    if b.BatterySize == 70 {
        rangeMiles = 240
    } else if b.BatterySize == 85 {
        rangeMiles = 270
    }
    fmt.Printf("This car can go approximately %d miles on a full charge.\n", rangeMiles)
}


//--

type ElectricCar struct {
    Car
    Battery Battery
}

func NewElectricCar(make, model string, year int) ElectricCar {
    return ElectricCar{
        Car: Car{
            Make:            make,
            Model:           model,
            Year:            year,
            OdometerReading: 20,
        },
        Battery: Battery{
            BatterySize: 70,
        },
    }
}

//--

func main() {
    myElectricCar := NewElectricCar("Tesla", "Model S", 2024)
    fmt.Println(myElectricCar.GetDescriptiveName())
    myElectricCar.ReadOdometer()
    myElectricCar.UpdateOdometer(15000)
    myElectricCar.IncrementOdometer(500)
    myElectricCar.ReadOdometer()
    myElectricCar.Battery.DescribeBattery()
    myElectricCar.Battery.GetRange()
}

// 2024 Tesla Model S
// This car has 20 miles on it.
// This car has 15500 miles on it.
// This car has a 70-kwh battery.
// This car can go approximately 240 miles on a full charge.
