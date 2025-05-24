package main

import (
	"fmt"
)

// Function to format an array of 10 strings
func FormatStringArray(arr [10]string) string {
	return fmt.Sprintf("(%s%s%s) %s%s%s-%s%s%s%s",
		arr[0], arr[1], arr[2], arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9])
}

func main() {
	// Example array of strings
	strArray := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	
	// Calling the function
	formatted := FormatStringArray(strArray)
	
	// Print result
	fmt.Println(formatted) // Output: (123) 456-7890
}
