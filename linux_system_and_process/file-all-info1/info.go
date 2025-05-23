package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func ExampleFileOperations() {
    filePath := "example.txt"
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Write data to the file
    data := []byte("Hello, World!")
    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    file.Close() // Close the file after writing to ensure data is saved

    // Read data from the file
    readData, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File content:", string(readData))

    // Get file information
    fileInfo, err := os.Stat(filePath)
    if err != nil {
        fmt.Println("Error getting file information:", err)
        return
    }
    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("File size:", fileInfo.Size())
    fmt.Println("File permissions:", fileInfo.Mode())

    // Rename the file
    newFilePath := "renamed.txt"
    err = os.Rename(filePath, newFilePath)
    if err != nil {
        fmt.Println("Error renaming file:", err)
        return
    }

    // Delete the file
    err = os.Remove(newFilePath)
    if err != nil {
        fmt.Println("Error deleting file:", err)
        return
    }

    // Get the current working directory
    workingDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting working directory:", err)
        return
    }
    fmt.Println("Current working directory:", workingDir)

    // List files in the current directory
    dirPath := "."
    files, err := os.ReadDir(dirPath)
    if err != nil {
        fmt.Println("Error listing files:", err)
        return
    }

    fmt.Println("Files in current directory:")
    for _, file := range files {
        fmt.Println(file.Name())
    }

    // Get the absolute path of a file
    absolutePath, err := filepath.Abs("example.txt")
    if err != nil {
        fmt.Println("Error getting absolute path:", err)
        return
    }
    fmt.Println("Absolute path:", absolutePath)

    // Create a buffered writer
    file, err = os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    writer := bufio.NewWriter(file)

    // Write data to the buffered writer
    data = []byte("Hello, World!\n")
    _, err = writer.Write(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // Flush the buffered writer to ensure data is written to the file
    err = writer.Flush()
    if err != nil {
        fmt.Println("Error flushing writer:", err)
        return
    }

    // Reopen the file for reading
    file, err = os.Open("output.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a buffered reader
    reader := bufio.NewReader(file)

    // Read data from the buffered reader
    readString, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        fmt.Println("Error reading from file:", err)
        return
    }

    fmt.Println("Read data:", readString)
}

func main() {
    ExampleFileOperations()
}


// go run info.go                                                                                                                                           
// File content: Hello, World!
// File name: example.txt
// File size: 13
// File permissions: -rw-rw-r--
// Current working directory: /home/mehanic/structure/system_i_process/file/info
// Files in current directory:
// info.go
// Absolute path: /home/mehanic/structure/system_i_process/file/info/example.txt
// Read data: Hello, World!
