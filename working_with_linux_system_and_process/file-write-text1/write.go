package main

import (
    "fmt"
    "os"
)

func main() {
    // Create a file called user.txt
    file, err := os.Create("user.txt")
    if err != nil {
        fmt.Println("error creating file:", err)
        return
    }
    defer file.Close()

    // Take user input
    var userinput string
    fmt.Print("Enter a word >>> ")
    fmt.Scan(&userinput)

    // Write user input to the file
    _, err = file.WriteString(userinput)
    if err != nil {
        fmt.Println("error writing to file:", err)
        return
    }
    file.Close() // Close the file after writing

    // Read the file
    fi, err := os.ReadFile("user.txt")
    if err != nil {
        fmt.Println("error reading file:", err)
        return
    }

    // Display the content of the file to the terminal
    fmt.Printf("file content: %s\n", fi)

    // Open a new file in append|create|readwrite mode called text.txt
    file, err = os.OpenFile("text.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("error opening file:", err)
        return
    }
    defer file.Close()

    // Declare variable and assign value to it
    line := "hello world\n"

    // Append variable to the file using Fprintln
    n, err := fmt.Fprintln(file, line)
    if err != nil {
        fmt.Println("error writing to file:", err)
        return
    }

    // Display how many bytes appended to file
    fmt.Println("Added successfully: total bytes", n)
}


// go run write.go                                                                                                                                           1 ↵ ──(Sat,Jan18)─┘
// Enter a word >>> terraform
// file content: terraform
// Added successfully: total bytes 13
