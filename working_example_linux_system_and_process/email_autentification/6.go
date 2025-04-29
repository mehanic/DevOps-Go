package main

import (
    "fmt"
)

func main() {
    var email string
    // Read input from the user
    fmt.Print("Enter your email: ")
    fmt.Scanf("%s", &email)

    isCorrect := false
    // Iterate over each character in the email string
    for i := 0; i < len(email); i++ {
        // Check if the character is '@'
        if email[i] == '@' {
            isCorrect = true
            break
        }
    }

    // Print appropriate message based on the presence of '@'
    if isCorrect {
        fmt.Println("ok")
    } else {
        fmt.Println("should contain @")
    }
}