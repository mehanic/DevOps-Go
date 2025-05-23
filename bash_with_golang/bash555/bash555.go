package main

import (
    "fmt"
    "io/ioutil"
    "log"
//    "os"
)

func main() {
    // Countdown from 6 to 1
    count := 6
    for count > 0 {
        fmt.Printf("Value of count is: %d\n", count)
        count-- // Decrement the counter
    }

    // Directory to list
    dir := "/bin"

    // Read files in the directory
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatalf("Failed to read directory %s: %v", dir, err)
    }

    // Print full path of each file
    for _, file := range files {
        fmt.Printf("File: %s/%s\n", dir, file.Name())
    }
}

