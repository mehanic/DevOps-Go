package main

import (
    "fmt"
    "strings"
)

type Car struct {
    Make             string
    Model            string
    Year             int
    OdometerReading  int
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

//--

type Battery struct {
    BatterySize int
}

func (b *Battery) DescribeBattery() {
    fmt.Printf("This car has a %d-kwh battery.\n", b.BatterySize)
}

func (b *Battery) GetRange() {
    var rangeMiles int
    switch b.BatterySize {
    case 70:
        rangeMiles = 240
    case 85:
        rangeMiles = 270
    default:
        rangeMiles = 0
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

func (e *ElectricCar) FillGasTank() {
    fmt.Println("This car doesn't need a gas tank!")
}

//---

func main() {
    // Create an ElectricCar instance
    myTesla := NewElectricCar("Tesla", "Model S", 2016)
    fmt.Println(myTesla.GetDescriptiveName())
    myTesla.Battery.DescribeBattery()
    myTesla.Battery.GetRange()
    myTesla.FillGasTank()

    // Create a Car instance
    myUsedCar := Car{Make: "Subaru", Model: "Outback", Year: 2013}
    fmt.Println(myUsedCar.GetDescriptiveName())
    myUsedCar.UpdateOdometer(23500)
    myUsedCar.ReadOdometer()
    myUsedCar.IncrementOdometer(100)
    myUsedCar.ReadOdometer()

    // Create another Car instance
    myNewCar := Car{Make: "Audi", Model: "A4", Year: 2016}
    fmt.Println(myNewCar.GetDescriptiveName())
    myNewCar.UpdateOdometer(25)
    myNewCar.OdometerReading = 23
    myNewCar.ReadOdometer()
}

// 2016 Tesla Model S
// This car has a 70-kwh battery.
// This car can go approximately 240 miles on a full charge.
// This car doesn't need a gas tank!
// 2013 Subaru Outback
// This car has 23500 miles on it.
// This car has 23600 miles on it.
// 2016 Audi A4
// This car has 23 miles on it.
