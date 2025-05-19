package main

import (
	"fmt"
	"runtime"
)

func func1(callerName string, source string, lineNum int) {
	// Print the function names and BASH_SOURCE equivalents
	fmt.Printf("func1: FUNCNAME0 is %s\n", "func1")
	fmt.Printf("func1: FUNCNAME1 is %s\n", callerName)
	fmt.Printf("func1: FUNCNAME2 is %s\n", "main")
	fmt.Printf("func1: BASH_SOURCE0 is %s\n", source)
	fmt.Printf("func1: BASH_SOURCE1 is %s\n", "main.go") // This is hardcoded since Go doesn't have an equivalent
	fmt.Printf("func1: BASH_SOURCE2 is %s\n", "") // No further BASH_SOURCE equivalent
	fmt.Printf("func1: LINENO is %d\n", lineNum)
	func2(callerName, source, lineNum+1) // Call func2, simulating line number increment
}

func func2(callerName string, source string, lineNum int) {
	// Similar structure for func2
	fmt.Printf("func2: FUNCNAME0 is %s\n", "func2")
	fmt.Printf("func2: FUNCNAME1 is %s\n", callerName)
	fmt.Printf("func2: FUNCNAME2 is %s\n", "main")
	fmt.Printf("func2: BASH_SOURCE0 is %s\n", source)
	fmt.Printf("func2: BASH_SOURCE1 is %s\n", "main.go") // This is hardcoded
	fmt.Printf("func2: BASH_SOURCE2 is %s\n", "") // No further BASH_SOURCE equivalent
	fmt.Printf("func2: LINENO is %d\n", lineNum)
}

func main() {
	// Get the caller function name and line number
	_, _, lineNum, _ := runtime.Caller(0)
	func1("main", "main.go", lineNum) // Call func1 with caller information
}

