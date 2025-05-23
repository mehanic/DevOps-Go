package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    // Check if exactly 3 command-line arguments are provided (including the script name)
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run script.go <user_name> <fave_color>")
        os.Exit(1)
    }

    // Assign command-line arguments to variables
    script := os.Args[0]
    userName := os.Args[1]
    faveColor := os.Args[2]
    prompt := "8====D~~~ "

    // Print initial messages
    fmt.Printf("Hi %s, I'm the %s script.\n", userName, script)
    fmt.Println("I'd like to ask you a few questions.")

    // Create a new scanner for reading input
    scanner := bufio.NewScanner(os.Stdin)

    // Ask questions and read user input
    fmt.Printf("Do you like me %s? %s", userName, prompt)
    scanner.Scan()
    likes := scanner.Text()

    fmt.Printf("Where do you live %s? %s", userName, prompt)
    scanner.Scan()
    lives := scanner.Text()

    fmt.Printf("What kind of computer do you have? %s", prompt)
    scanner.Scan()
    computer := scanner.Text()

    // Print the results
    fmt.Printf(`
Alright, so you said %q about liking me.
You live in %q. Not sure where that is.
And you have a %q computer. Nice.
Your favorite color is %q. Mine too!
`, likes, lives, computer, faveColor)
}

