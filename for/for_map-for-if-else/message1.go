package main

import (
	"bufio"
	"fmt"
	"os"
//	"strings"
)

func main() {
	birthdays := map[string]string{
		"Alice": "Apr 1",
		"Bob":   "Dec 12",
		"Carol": "Mar 4",
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a name: (blank to quit)\n")
		scanner.Scan()
		name := scanner.Text()
		if name == "" {
			break
		}

		if bday, ok := birthdays[name]; ok {
			fmt.Printf("%s is the birthday of %s\n", bday, name)
		} else {
			fmt.Printf("I do not have birthday information for %s\n", name)
			fmt.Print("What is their birthday?\n")
			scanner.Scan()
			bday := scanner.Text()
			birthdays[name] = bday
			fmt.Println("Birthday database updated.")
		}
	}
}

