package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	timeLeft := 10
	for timeLeft > 0 {
		fmt.Print(timeLeft)
		time.Sleep(1 * time.Second)
		timeLeft--
	}

	// At the end of the countdown, play a sound file.
	err := exec.Command("cmd", "/c", "start", "alarm.wav").Run()
	if err != nil {
		fmt.Println("Error playing sound:", err)
	}
}

// 10987654321Error playing sound: exec: "cmd": executable file not found in $PATH

