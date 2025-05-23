package main

import "fmt"

func main() {
    // sets the variable "people" to 50
    people := 50

    // sets the variable "cars" to 60
    cars := 60

    // sets the variable "buses" to 50
    buses := 50

    // if "cars is greater than people"
    if cars > people {
        // print string
        fmt.Println("We should take the cars.")
    } else if cars < people {
        // else if "cars is less than people"
        fmt.Println("We should not take the cars.")
    } else {
        // else
        fmt.Println("We can't decide.")
    }

    // if "buses greater than cars"
    if buses > cars {
        // print string
        fmt.Println("That's too many buses.")
    } else if buses < cars {
        // else if "buses less than cars"
        fmt.Println("Maybe we could take the buses.")
    } else {
        // else
        fmt.Println("We still can't decide.")
    }

    // if "people is greater than buses"
    if people > buses {
        // print string
        fmt.Println("Alright, let's just take the buses.")
    } else {
        // else
        fmt.Println("Fine, let's stay home then.")
    }
}

// We should take the cars.
// Maybe we could take the buses.
// Fine, let's stay home then.
