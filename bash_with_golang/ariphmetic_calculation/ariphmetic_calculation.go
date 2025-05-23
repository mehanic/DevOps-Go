package main

import (
	"fmt"
)

func main() {
	no1 := 11
	no2 := 22

	no1 += 6  // equivalent to let no1+=6 in Bash
	no2 -= 5  // equivalent to let no2-=5 in Bash

	result1 := no1 + no2           // $[ no1 + no2 ] in Bash
	result2 := no1 + no2           // $[ $no1 + no2 ] in Bash
	result3 := 3 + 4               // equivalent to expr 3 + 4
	result4 := no1 + 5             // equivalent to expr $no1 + 5

	// Output results
	fmt.Printf("result1 (no1 + no2): %d\n", result1)
	fmt.Printf("result2 (no1 + no2): %d\n", result2)
	fmt.Printf("result3 (3 + 4): %d\n", result3)
	fmt.Printf("result4 (no1 + 5): %d\n", result4)

	no := 54.0                     // Change to float64 for decimal calculations
	resultFloat := no * 1.5        // equivalent to echo "$no * 1.5" | bc in Bash
	fmt.Printf("resultFloat (no * 1.5): %.2f\n", resultFloat)

	resultCalc := 4 * 0.56         // This is also a float calculation
	fmt.Printf("4 * 0.56: %.2f\n", resultCalc)
}

