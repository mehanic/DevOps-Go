package main

import (
	"fmt"
)

func main() {
	// Declare the strings
	str1 := "USA"
	str2 := "Canada"

	// Compare str1 and str2 (equivalent to `[ $str1 = $str2 ]`)
	if str1 == str2 {
		fmt.Println(0) // True in Bash would return 0
	} else {
		fmt.Println(1) // False in Bash would return 1
	}

	// Debug mode in Bash (`set -x`), we're just printing that debug mode would start
	fmt.Println("Debug mode starts (set -x equivalent)")

	// Compare str1 and str2 for inequality (equivalent to `[ $str1 != $str2 ]`)
	if str1 != str2 {
		fmt.Println(0) // True, returns 0
	} else {
		fmt.Println(1) // False, returns 1
	}

	// Check if str1 is an empty string (equivalent to `[ -z $str1 ]`)
	if len(str1) == 0 {
		fmt.Println(0) // True, returns 0
	} else {
		fmt.Println(1) // False, returns 1
	}

	// End debug mode in Bash (`set +x`), we're just printing that debug mode would stop
	fmt.Println("Debug mode ends (set +x equivalent)")

	// Check if str2 is non-empty (equivalent to `[ -n $str2 ]`)
	if len(str2) > 0 {
		fmt.Println(0) // True, returns 0
	} else {
		fmt.Println(1) // False, returns 1
	}

	// Exit (equivalent to `exit 0` in Bash)
	fmt.Println("Exit 0")
}

