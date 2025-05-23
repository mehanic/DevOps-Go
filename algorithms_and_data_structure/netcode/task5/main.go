// Spreadsheets often use this alphabetical encoding for its columns: "A", "B", "C", ..., "AA", "AB", ..., "ZZ", "AAA", "AAB", ....

// Given a column number, return its alphabetical column id. For example, given 1, return "A". Given 27, return "AA".
package main

import "fmt"

func columnId(columnNumber int) string {
	// Edge case: column number must be positive
	if columnNumber <= 0 {
	  return ""
	}
  
	// Initialize the result string
	result := ""
  
	// Iterate until columnNumber becomes 0
	for columnNumber > 0 {
	  // Get the remainder when columnNumber is divided by 26
	  remainder := columnNumber % 26
  
	  // If the remainder is 0, set it to 26 and decrement columnNumber by 1
	  if remainder == 0 {
		remainder = 26
		columnNumber -= 1
	  }
  
	  // Append the corresponding letter to the result string
	  result = string(remainder + 'A' - 1) + result
  
	  // Divide columnNumber by 26
	  columnNumber /= 26
	}
  
	return result
  }
  
  func main() {
	testCases := []int{1, 2, 3, 26, 27, 28, 52, 53, 702, 703, 704, 18278}

	for _, num := range testCases {
		fmt.Printf("Column Number: %d => Column Name: %s\n", num, columnId(num))
	}
}