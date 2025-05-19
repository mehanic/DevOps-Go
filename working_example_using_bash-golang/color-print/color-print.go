package main

import "fmt"

// Define ANSI escape codes for text formatting
const (
	txtUnd  = "\033[4m"   // Underline
	txtBld  = "\033[1m"   // Bold
	txtRed  = "\033[31m"  // Red
	txtGrn  = "\033[32m"  // Green
	txtYlw  = "\033[33m"  // Yellow
	txtBlu  = "\033[34m"  // Blue
	txtPur  = "\033[35m"  // Purple
	txtCyn  = "\033[36m"  // Cyan
	txtWht  = "\033[37m"  // White
	txtrst  = "\033[0m"   // Reset
)

func main() {
	// Use the defined constants to format output
	fmt.Printf("%sThis is bold text output from Go program%s\n", txtBld, txtrst)
	fmt.Printf("%sThis is coloured red except %sthis is blue%s\n", txtRed, txtBlu, txtrst)

	// Reset the terminal formatting
	fmt.Print(txtrst)
}

