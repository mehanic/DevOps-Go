package main

import "fmt"

func main() {
    name := "Zed A. Shaw"
    age := 35
    height := 74 // inches
    heightInCm := float64(height) * 2.54
    weight := 180 // lbs
    weightInKg := float64(weight) * 0.45359237
    eyes := "Blue"
    teeth := "White"
    hair := "Brown"

    fmt.Printf("Let's talk about %s.\n", name)
    fmt.Printf("He's %d inches tall (%.2f cm).\n", height, heightInCm)
    fmt.Printf("He's %d pounds heavy (%.2f kg).\n", weight, weightInKg)
    fmt.Println("Actually that's not too heavy.")
    fmt.Printf("He's got %s eyes and %s hair.\n", eyes, hair)
    fmt.Printf("His teeth are usually %s depending on the coffee.\n", teeth)

    // This line is tricky, try to get it exactly right
    fmt.Printf("If I add %d, %d, and %d I get %d.\n", age, height, weight, age + height + weight)
}

// Let's talk about Zed A. Shaw.
// He's 74 inches tall (187.96 cm).
// He's 180 pounds heavy (81.65 kg).
// Actually that's not too heavy.
// He's got Blue eyes and Brown hair.
// His teeth are usually White depending on the coffee.
// If I add 35, 74, and 180 I get 289.
