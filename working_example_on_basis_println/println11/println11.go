package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // Prompt the user for their age
    fmt.Print("How old are you? ")
    age, _ := reader.ReadString('\n')
    age = age[:len(age)-1] // Remove the newline character

    // Prompt the user for their height
    fmt.Print("How tall are you? ")
    height, _ := reader.ReadString('\n')
    height = height[:len(height)-1] // Remove the newline character

    // Prompt the user for their weight
    fmt.Print("How much do you weigh? ")
    weight, _ := reader.ReadString('\n')
    weight = weight[:len(weight)-1] // Remove the newline character

    // Print the collected information
    fmt.Printf("So, you're %s old, %s tall and %s heavy.\n", age, height, weight)
}

// How old are you? 30
// How tall are you? 140
// How much do you weigh? 56
// So, you're 30 old, 140 tall and 56 heavy.
