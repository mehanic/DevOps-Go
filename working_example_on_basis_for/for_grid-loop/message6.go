package main

import "fmt"

// Initialize the grid
var grid = [][]string{
    {".", ".", ".", ".", ".", "."},
    {".", "0", "0", ".", ".", "."},
    {"0", "0", "0", "0", ".", "."},
    {"0", "0", "0", "0", "0", "."},
    {".", "0", "0", "0", "0", "0"},
    {"0", "0", "0", "0", "0", "."},
    {"0", "0", "0", "0", ".", "."},
    {".", "0", "0", ".", ".", "."},
    {".", ".", ".", ".", ".", "."},
}

// Function to print the grid in the desired order
func makeGrid() {
    for i := 0; i < 6; i++ { // Loop over columns
        for j := 0; j < 9; j++ { // Loop over rows
            fmt.Print(grid[j][i]) // Print each element without a newline
        }
        fmt.Println() // Newline after each row
    }
}

func main() {
    makeGrid()
}

