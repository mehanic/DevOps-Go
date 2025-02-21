package main

import "fmt"

func main() {
    // We have 100 cars
    cars := 100

    // You can fit 4 people in a car
    spaceInACar := 4

    // We have 30 drivers
    drivers := 30

    // We have 90 passengers
    passengers := 90

    // After all of the drivers are in cars, we have "cars not driven" left over
    carsNotDriven := cars - drivers

    // The number of cars driven equals the number of drivers
    carsDriven := drivers

    // Carpool capacity is the number of cars driven multiplied by the space in each car
    carpoolCapacity := float64(carsDriven) * float64(spaceInACar)

    // The average number of passengers in each car equals the number of passengers divided by the number of cars available
    averagePassengersPerCar := float64(passengers) / float64(carsDriven)

    fmt.Println("There are", cars, "cars available.")
    fmt.Println("There are only", drivers, "drivers available.")
    fmt.Println("There will be", carsNotDriven, "empty cars today.")
    fmt.Println("We can transport", int(carpoolCapacity), "people today.")
    fmt.Println("We have", passengers, "to carpool today.")
    fmt.Println("We need to put about", averagePassengersPerCar, "in each car.")

    // Extra credit: Example calculation
    i := 40
    x := 234
    j := 23.4
    result := float64(i) * float64(x) / j
    fmt.Printf("Extra credit calculation result: %.1f\n", result)
}

// There are 100 cars available.
// There are only 30 drivers available.
// There will be 70 empty cars today.
// We can transport 120 people today.
// We have 90 to carpool today.
// We need to put about 3 in each car.
// Extra credit calculation result: 400.0
