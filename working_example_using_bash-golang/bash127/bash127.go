package main

import (
    "encoding/base64"
    "flag"
    "fmt"
    "log"
    "math/rand"
    "time"
)

func generatePassword(length int) (string, error) {
    // Generate 48 random bytes
    randomBytes := make([]byte, 48)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }

    // Encode the random bytes to base64
    encoded := base64.StdEncoding.EncodeToString(randomBytes)

    // Return the substring of the specified length
    if length > len(encoded) {
        length = len(encoded) // Prevent out of bounds
    }

    return encoded[:length], nil
}

func main() {
    var numberOfPasswords int
    var passwordLength int

    // Parse command-line arguments
    flag.IntVar(&numberOfPasswords, "n", 1, "Number of passwords to generate")
    flag.IntVar(&passwordLength, "l", 64, "Length of each password")
    flag.Parse()

    // Set defaults if arguments are not provided
    if numberOfPasswords <= 0 {
        numberOfPasswords = 1
    }
    if passwordLength <= 0 {
        passwordLength = 64
    }

    // Seed the random number generator (optional)
    rand.Seed(time.Now().UnixNano())

    // Generate the passwords
    for i := 0; i < numberOfPasswords; i++ {
        password, err := generatePassword(passwordLength)
        if err != nil {
            log.Fatalf("Error generating password: %v", err)
        }
        fmt.Println(password)
    }
}

