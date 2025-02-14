package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Itoa doesn't return any errors
	// So, you don't have to handle the errors for it

	s := strconv.Itoa(11)

	fmt.Println(s)
	main1()
	main2()
}

func main1() {
	// Atoi returns an error value
	// So, you should always check it

	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}

func main2() {
	age := os.Args[1]

	// Atoi returns an int and an error value
	// So, you need to handle the errors

	n, err := strconv.Atoi(age)

	// handle the error immediately and quit
	// don't do anything else here

	if err != nil {
		fmt.Println("ERROR:", err)

		// quits/returns from the main function
		// so, the program ends
		return
	}

	// DO NOT DO THIS:
	// else {
	//   happy path
	// }

	// DO THIS:

	// happy path (err is nil)
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}


// 11
// panic: runtime error: index out of range [1] with length 1

// goroutine 1 [running]:
// main.main1()
// 	/home/mehanic/structure/13.error/error-handling/error-handling.go:24 +0x11b
// main.main()
// 	/home/mehanic/structure/13.error/error-handling/error-handling.go:16 +0x5f
// exit status 2

