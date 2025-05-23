package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var wakeHour, wakeMinute int

	// Get user input for wake-up time
	fmt.Print("What time do you Wake Up? (Hour): ")
	fmt.Scan(&wakeHour)
	fmt.Print("and Minutes: ")
	fmt.Scan(&wakeMinute)

	// Get current time
	now := time.Now()
	curHour := now.Hour()
	curMinute := now.Minute()

	fmt.Printf("You Selected: %02d:%02d\n", wakeHour, wakeMinute)
	fmt.Printf("It is Currently: %02d:%02d\n", curHour, curMinute)

	var hoursLeft, minutesLeft int

	// Calculate hours left
	if curHour < wakeHour {
		hoursLeft = wakeHour - curHour
	} else if curHour > wakeHour {
		hoursLeft = curHour - wakeHour
	} else {
		hoursLeft = 0
	}

	// Calculate minutes left
	if curMinute < wakeMinute {
		minutesLeft = wakeMinute - curMinute
	} else if curMinute > wakeMinute {
		minutesLeft = curMinute - wakeMinute
	} else {
		minutesLeft = 0
	}

	if hoursLeft == 0 && minutesLeft == 0 {
		fmt.Println("Taking a nap?")
	} else {
		if hoursLeft > 0 {
			fmt.Printf("%d hours left\n", hoursLeft)
		} else {
			fmt.Printf("%d minutes left\n", minutesLeft)
		}
	}

	fmt.Printf("Sleeping for %d hours and %d minutes\n", hoursLeft, minutesLeft)

	// Sleep for the calculated time
	time.Sleep(time.Duration(hoursLeft) * time.Hour)
	time.Sleep(time.Duration(minutesLeft) * time.Minute)

	// Play the alarm sound
	err := exec.Command("mplayer", "/home/username/.alarm/alarm.mp3").Run() // Update the path as needed
	if err != nil {
		fmt.Printf("Error playing alarm: %v\n", err)
	}
}

