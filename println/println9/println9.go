package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    // Prompt the user for their name
    fmt.Print("What's your name?\n==> ")
    name, _ := reader.ReadString('\n')
    name = name[:len(name)-1] // Remove the newline character

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
    fmt.Printf("So, your name is %s, you're %s old, %s tall and %s heavy.\n",
        name, age, height, weight)
}

// What's your name?
// ==> max
// How old are you? 30
// How tall are you? 160
// How much do you weigh? 45
// So, your name is max, you're 30 old, 160 tall and 45 heavy.
