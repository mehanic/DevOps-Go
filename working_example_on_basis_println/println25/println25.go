package main

import "fmt"

func main() {
    people := 20
    cats := 15
    dogs := 30

    if people < cats {
        fmt.Println("Too many cats! The world is doomed!")
    }

    if people > cats {
        fmt.Println("Not many cats! The world is saved!")
    }

    if people < dogs {
        fmt.Println("The world is drooled on!")
    }

    if people > dogs {
        fmt.Println("The world is dry!")
    }

    dogs += 5

    if people >= dogs {
        fmt.Println("People are greater than or equal to dogs.")
    }

    if people <= dogs {
        fmt.Println("People are less than or equal to dogs.")
    }

    if people == dogs {
        fmt.Println("People are dogs.")
    }
}


// Not many cats! The world is saved!
// The world is drooled on!
// People are less than or equal to dogs.
