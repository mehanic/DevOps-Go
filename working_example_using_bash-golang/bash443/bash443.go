package main

import (
	"fmt"
	"os/exec"
//	"strconv"
	"time"
)

func main() {
	var hour, minute int

	// Get wake-up time from the user
	fmt.Print("What time do you wake up? (Hour): ")
	fmt.Scan(&hour)
	fmt.Print("And Minutes? ")
	fmt.Scan(&minute)

	// Get current time
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()

	// Display the user input and current time
	fmt.Printf("You selected %02d:%02d\n", hour, minute)
	fmt.Printf("\nIt is currently %02d:%02d\n", currentHour, currentMinute)

	// Calculate the time difference
	var hourLeft, minLeft int
	if currentHour < hour {
		hourLeft = hour - currentHour
	} else if currentHour > hour {
		hourLeft = currentHour - hour
		fmt.Printf("\n%d - %d means you have %d hours left\n", currentHour, hour, hourLeft)
	} else {
		hourLeft = 0
		fmt.Println("Taking a nap?")
	}

	if currentMinute < minute {
		minLeft = minute - currentMinute
		fmt.Printf("%d - %d means you have %d minutes still\n", minute, currentMinute, minLeft)
	} else if currentMinute > minute {
		minLeft = currentMinute - minute
		fmt.Printf("%d - %d means you have %d minutes left\n", currentMinute, minute, minLeft)
	} else {
		minLeft = 0
		fmt.Println("And no minutes")
	}

	// Sleep for the remaining hours and minutes
	fmt.Printf("Sleeping for %d hours and %d minutes\n", hourLeft, minLeft)
	time.Sleep(time.Duration(hourLeft) * time.Hour)
	time.Sleep(time.Duration(minLeft) * time.Minute)

	// Play the alarm sound using mplayer
	fmt.Println("Wake up!")
	cmd := exec.Command("mplayer", "~/.alarm/alarm.mp3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing alarm:", err)
	}
}

