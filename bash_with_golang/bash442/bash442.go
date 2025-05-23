package main

import (
	"fmt"
	"os"
	"os/exec"
//	"strconv"
	"time"
)

func main() {
	var h, m int

	// Prompt user for wake-up time
	fmt.Print("What time do you Wake Up? (HH): ")
	_, err := fmt.Scanf("%d", &h)
	if err != nil {
		fmt.Println("Invalid input for hours:", err)
		return
	}

	fmt.Print("and Minutes? (MM): ")
	_, err = fmt.Scanf("%d", &m)
	if err != nil {
		fmt.Println("Invalid input for minutes:", err)
		return
	}

	// Get the current time
	currentTime := time.Now()
	curH := currentTime.Hour()
	curM := currentTime.Minute()

	fmt.Printf("You Selected: %02d:%02d\n", h, m)
	fmt.Printf("It is Currently: %02d:%02d\n", curH, curM)

	hourLeft, minLeft := 0, 0

	// Calculate hours left
	if curH < h {
		hourLeft = h - curH
		fmt.Printf("%d - %d means You Have: %d hours still\n", h, curH, hourLeft)
	} else if curH > h {
		hourLeft = curH - h
		fmt.Printf("%d - %d means you have %d hours left\n", curH, h, hourLeft)
	} else {
		hourLeft = 0
		fmt.Println("Taking a nap?")
	}

	// Calculate minutes left
	if curM < m {
		minLeft = m - curM
		fmt.Printf("%d - %d means you have: %d minutes still\n", m, curM, minLeft)
	} else if curM > m {
		minLeft = curM - m
		fmt.Printf("%d - %d means you have %d minutes left\n", curM, m, minLeft)
	} else {
		minLeft = 0
		fmt.Println("and no minutes\n")
	}

	// Combine hours and minutes for total wait time
	totalDuration := time.Duration(hourLeft)*time.Hour + time.Duration(minLeft)*time.Minute
	fmt.Printf("Sleeping for %d hours and %d minutes\n", hourLeft, minLeft)

	// Sleep for the calculated duration
	time.Sleep(totalDuration)

	// Play alarm sound
	err = playAlarm()
	if err != nil {
		fmt.Println("Error playing alarm:", err)
	}
}

func playAlarm() error {
	// Adjust the command below according to your OS and alarm file path
	cmd := exec.Command("mplayer", "~/.alarm/alarm.mp3")
	cmd.Dir = os.Getenv("HOME") // Ensure it runs in the user's home directory
	return cmd.Run()
}

