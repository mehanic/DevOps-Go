package main

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	// Set up the prompt
	prompt := promptui.Prompt{
		Label: "Timer",
		Templates: &promptui.PromptTemplates{
			Prompt:   "{{ . | green }}", // Prompt text color
			Success:  "{{ . | faint }}", // Success message color
			Failure:  "{{ . | red }}",   // Failure message color
		},
		Help: "Enter a duration for the timer.\n\n Use 10s for 10 seconds, 20m for 20 minutes, or 3h for 3 hours.",
	}

	// Ask for user input
	timeInput, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Parse the duration
	duration, err := time.ParseDuration(timeInput)
	if err != nil {
		fmt.Printf("Invalid duration: %v\n", err)
		return
	}

	// Sleep for the specified duration
	time.Sleep(duration)

	// Show timer complete message
	fmt.Printf("The timer is over.\n\n It has been %s.\n", timeInput)
}

