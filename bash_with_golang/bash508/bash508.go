package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// checkArguments validates if the exact number of expected arguments are provided.
func checkArguments(expected int, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("less than 1 argument received, exiting")
	}
	if len(args) != expected {
		return fmt.Errorf("incorrect number of arguments, expected %d but got %d", expected, len(args))
	}
	return nil
}

// checkInteger checks if the given argument is a valid integer.
func checkInteger(arg string) error {
	_, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("the argument is not a valid integer: %s", arg)
	}
	return nil
}

// checkYesNo checks if the user response is "yes" or "no".
func checkYesNo(response string) (bool, error) {
	lowerResponse := strings.ToLower(response)
	if lowerResponse == "y" || lowerResponse == "yes" {
		return true, nil
	} else if lowerResponse == "n" || lowerResponse == "no" {
		return false, nil
	}
	return false, fmt.Errorf("neither yes nor no, exiting")
}

// setCwd sets the current working directory to the script's location.
func setCwd() error {
	executable, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}
	dir := filepath.Dir(executable)
	return os.Chdir(dir)
}

func main() {
	// Example usages of the functions:

	// 1. Argument check (assume expecting exactly 2 arguments)
	args := os.Args[1:] // Exclude the program name itself
	if err := checkArguments(2, args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Arguments check passed")

	// 2. Integer check (assume we expect the first argument to be an integer)
	if err := checkInteger(args[0]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Integer check passed")

	// 3. Yes/No check (assume we expect the second argument to be yes or no)
	if _, err := checkYesNo(args[1]); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Yes/No check passed")

	// 4. Set current working directory to script location
	if err := setCwd(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Changed working directory to script location")
}

