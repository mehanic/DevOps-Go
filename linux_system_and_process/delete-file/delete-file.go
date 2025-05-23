package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func main() {
    var reqPath string
    fmt.Print("Enter your path: ")
    fmt.Scanln(&reqPath)

    age := 3

    if _, err := os.Stat(reqPath); os.IsNotExist(err) {
        fmt.Println("Please provide valid path")
        os.Exit(1)
    }

    fileInfo, err := os.Stat(reqPath)
    if err != nil {
        fmt.Println("Error accessing the path:", err)
        os.Exit(1)
    }

    if !fileInfo.IsDir() {
        fmt.Println("Please provide directory path")
        os.Exit(2)
    }

    today := time.Now()

    files, err := os.ReadDir(reqPath)
    if err != nil {
        fmt.Println("Error reading the directory:", err)
        os.Exit(1)
    }

    for _, file := range files {
        eachFilePath := filepath.Join(reqPath, file.Name())

        fileInfo, err := os.Stat(eachFilePath)
        if err != nil {
            fmt.Println("Error accessing the file:", err)
            continue
        }

        if fileInfo.Mode().IsRegular() {
            fileCreDate := fileInfo.ModTime()
            difDays := int(today.Sub(fileCreDate).Hours() / 24)

            if difDays > age {
                fmt.Println(eachFilePath, difDays)
            }
        }
    }
}
