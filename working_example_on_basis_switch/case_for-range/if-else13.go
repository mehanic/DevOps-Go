package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

// Coordinates and colors
var (
	nameField           = robotgo.Point{X: 648, Y: 319}
	submitButton        = robotgo.Point{X: 651, Y: 817}
	submitButtonColor   = robotgo.CHex{R: 75, G: 141, B: 249}
	submitAnotherLink   = robotgo.Point{X: 760, Y: 224}
	formData            = []map[string]interface{}{
		{"name": "Alice", "fear": "eavesdroppers", "source": "wand",
			"robocop": 4, "comments": "Tell Bob I said hi."},
		{"name": "Bob", "fear": "bees", "source": "amulet", "robocop": 4,
			"comments": "n/a"},
	}
)
func main() {
	robotgo.SetDelay(500) // Set a default delay of 0.5 seconds

	for _, person := range formData {
		fmt.Println(">>> 5 SECOND PAUSE TO LET USER PRESS CTRL-C <<<")
		time.Sleep(5 * time.Second)

		for {
			if robotgo.GetPixelColor(submitButton.X, submitButton.Y) == submitButtonColor {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Printf("Entering %s info...\n", person["name"])
		robotgo.MoveClick(nameField.X, nameField.Y)
		robotgo.TypeStr(fmt.Sprintf("%s\t", person["name"]))
		robotgo.TypeStr(fmt.Sprintf("%s\t", person["fear"]))

		switch person["source"] {
		case "wand":
			robotgo.KeyTap("down")
		case "amulet":
			robotgo.KeyTap("down", "down")
		case "crystal ball":
			robotgo.KeyTap("down", "down", "down")
		case "money":
			robotgo.KeyTap("down", "down", "down", "down")
		}

		switch person["robocop"] {
		case 1:
			robotgo.KeyTap("space")
		case 2:
			robotgo.KeyTap("right")
		case 3:
			robotgo.KeyTap("right", "right")
		case 4:
			robotgo.KeyTap("right", "right", "right")
		case 5:
			robotgo.KeyTap("right", "right", "right", "right")
		}

		robotgo.TypeStr(fmt.Sprintf("%s\t", person["comments"]))
		robotgo.KeyTap("enter")
		fmt.Println("Clicked Submit.")
		time.Sleep(5 * time.Second)
		robotgo.MoveClick(submitAnotherLink.X, submitAnotherLink.Y)
	}
}

// The error message indicates that the X11/extensions/XTest.h header file is missing, 
// which is required by the robotgo library when compiling on Linux systems. robotgo relies on 
// several system-level dependencies to interact with the keyboard, mouse, and screen.

// sudo apt update
// sudo apt install -y libx11-dev libxtst-dev libpng-dev xorg-dev