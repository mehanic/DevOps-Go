package main

import (
	"fmt"
	"os/exec"
//	"strconv"
)

func main() {
	for {
		var choice string
		fmt.Print("1: Show disk usage.  2: Show uptime. ")

		// Read user input
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch choice {
		case "1":
			// Show disk usage
			cmd := exec.Command("df", "-h")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error executing df:", err)
				continue
			}
			fmt.Println(string(output))

		case "2":
			// Show uptime
			cmd := exec.Command("uptime")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error executing uptime:", err)
				continue
			}
			fmt.Println(string(output))

		default:
			// Exit the loop for any other input
			fmt.Println("Exiting...")
			return
		}
	}
}

