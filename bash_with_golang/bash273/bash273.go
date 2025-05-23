package main

import (
	"fmt"
	"runtime"
)

func func1(callerName string, lineNum int) {
	// Get the name of the calling function and the line number
	fmt.Printf("func1: FUNCNAME0 is %s\n", "func1")
	fmt.Printf("func1: FUNCNAME1 is %s\n", callerName)
	fmt.Printf("func1: FUNCNAME2 is %s\n", "main")
	fmt.Printf("func1: LINENO is %d\n", lineNum)
	func2("func1", lineNum+1) // Call func2, simulating the line number increment
}

func func2(callerName string, lineNum int) {
	// Get the name of the calling function and the line number
	fmt.Printf("func2: FUNCNAME0 is %s\n", "func2")
	fmt.Printf("func2: FUNCNAME1 is %s\n", callerName)
	fmt.Printf("func2: FUNCNAME2 is %s\n", "main")
	fmt.Printf("func2: LINENO is %d\n", lineNum)
}

func main() {
	// Get the caller function name and line number
	_, _, lineNum, _ := runtime.Caller(0)
	func1("main", lineNum)
}

