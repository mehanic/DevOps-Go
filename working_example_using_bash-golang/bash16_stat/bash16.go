package main

import (
	"fmt"
	"os"
	"os/user"
)

// Function to mimic my_function_global
func myFunctionGlobal() (int, error) {
	usr, err := user.Current()
	if err != nil {
		return 255, err
	}

	// Check if .bashrc exists
	_, err = os.Stat("/home/" + usr.Username + "/.bashrc")
	if os.IsNotExist(err) {
		return 1, nil // Indicate file does not exist
	}
	return 0, nil // Indicate success
}

// Function to mimic my_function_return
func myFunctionReturn() int {
	usr, _ := user.Current()
	_, err := os.Stat("/home/" + usr.Username + "/.bashrc")
	if os.IsNotExist(err) {
		return 1 // Indicate file does not exist
	}
	return 0 // Indicate success
}

// Function to mimic my_function_str
func myFunctionStr(uname string) string {
	_, err := os.Stat("/home/" + uname + "/.bashrc")
	if os.IsNotExist(err) {
		return "NOT FOUND"
	}
	return "FOUND IT"
}

func main() {
	globalRet := 255

	// Mimic the first function call
	fmt.Printf("Current ret: %d\n", globalRet)
	ret, err := myFunctionGlobal()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	globalRet = ret
	fmt.Printf("Current ret after: %d\n", globalRet)

	// Reset and call the second function
	globalRet = 255
	fmt.Printf("Current ret: %d\n", globalRet)
	globalRet = myFunctionReturn()
	fmt.Printf("Current ret after: %d\n", globalRet)

	// And for giggles, we can pass back output too!
	var output string
	fmt.Printf("Current ret: %s\n", output)
	usr, _ := user.Current()
	output = myFunctionStr(usr.Username) // Using the current user
	fmt.Printf("Current ret after: %s\n", output)
}
