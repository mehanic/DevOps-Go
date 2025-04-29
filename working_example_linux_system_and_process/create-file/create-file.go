package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    // Open a file for writing
    fo, err := os.Create("newdemo.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    fmt.Println("w") // Mode is write
    fmt.Println(false) // File is not readable in write mode
    fmt.Println(true)  // File is writable
    fo.Close()

    // Write to a file
    myContent := []string{"This is a data -1\n", "This is a data-2\n", "This is a data-3"}
    fo, err = os.Create("random.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    fo.WriteString("This is a first line\n")
    fo.WriteString("This is a second line\n")
    fo.WriteString("This is a third line")
    // Uncomment the following line to write the content of myContent
    // for _, line := range myContent {
    // 	fo.WriteString(line)
    // }
    fo.Close()

    // Write using a loop
    myContent = []string{"This is using loop-iteratioin-1", "This is using loop-iterantion-2", "This is using loop-iteratioin-3"}
    fo, err = os.OpenFile("with_loop.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    for _, eachLine := range myContent {
        fo.WriteString(eachLine + "\n")
    }
    fo.Close()

    // Read the entire file content
    fo, err = os.Open("with_loop.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    data, err := io.ReadAll(fo)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fo.Close()
    fmt.Println(string(data))

    // Read line by line
    fo, err = os.Open("with_loop.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    reader := bufio.NewReader(fo)
    line, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading line:", err)
        return
    }
    fmt.Print(line)
    line, err = reader.ReadString('\n')
    if err != nil && err != io.EOF {
        fmt.Println("Error reading line:", err)
        return
    }
    fmt.Print(line)
    fo.Close()

    // Read all lines into a slice
    fo, err = os.Open("with_loop.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    scanner := bufio.NewScanner(fo)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error scanning file:", err)
        return
    }
    fo.Close()

    // Print specific lines
    for i := 0; i < 3; i++ {
        fmt.Println(lines[i]) // lines[0], lines[1], lines[2]
    }

    // Print the last line
    if len(lines) > 0 {
        fmt.Println(lines[len(lines)-1])
    }
}


// w
// false
// true
// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-3
// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-3
// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-3
// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-3

// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-1
// This is using loop-iterantion-2
// This is using loop-iteratioin-3
// This is using loop-iteratioin-3
