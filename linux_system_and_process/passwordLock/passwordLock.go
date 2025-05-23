package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)
//passwords: A map that stores passwords for different accounts. The keys are account 
//names (e.g., "email", "blog"), and the values are the corresponding passwords.
func main() {
	passwords := map[string]string{
		"email":   "ihwbrfh3rhbfwhbdifuhsud",
		"blog":    "974y5h34kr73ihfi3",
		"luggage": "12345",
	}
//	os.Args: A slice that contains the command-line arguments passed to the program.
//	This block checks if the user has provided an account name as a command-line argument. 
//	If not, it prints a usage message 
//	and exits the program with an error code (1).
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./pw [account] - copy account password.")
		os.Exit(1)
	}

	account := os.Args[1]

	if password, ok := passwords[account]; ok {
		clipboard.WriteAll(password)
		fmt.Println("The password for", account, "has been copied to the clipboard.")
	} else {
		fmt.Println("There is no account named", account)
	}

	os.Exit(0)
}


// go run passwordLock.go email                                                                                                                              ──(Sun,Jan19)─┘
// The password for email has been copied to the clipboard.
