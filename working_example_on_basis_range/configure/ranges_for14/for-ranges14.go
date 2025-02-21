package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to read user input from the console
func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func goldRoom() {
	fmt.Println("This room is full of gold. How much do you take?")

	for {
		next := getInput("==> ")
		howMuch, err := strconv.Atoi(next)
		if err != nil {
			fmt.Println("That's not a number... try again!")
			continue
		}

		if howMuch < 50 {
			fmt.Println("Nice, you're not greedy, you win!")
			os.Exit(0)
		} else {
			dead("You greedy bastard!")
		}
	}
}

func bearRoom() {
	fmt.Println("There is a bear here.")
	fmt.Println("The bear has a bunch of honey.")
	fmt.Println("The fat bear is in front of another door.")
	fmt.Println("How are you going to move the bear?")
	bearMoved := false

	for {
		next := getInput("==> ")

		switch {
		case next == "take honey":
			dead("The bear looks at you and then pimp slaps your face off.")
		case next == "taunt bear" && !bearMoved:
			fmt.Println("The bear has moved from the door. You can go through it now.")
			bearMoved = true
		case next == "taunt bear" && bearMoved:
			dead("The bear gets pissed off and chews your crotch off.")
		case next == "open door" && bearMoved:
			goldRoom()
		default:
			fmt.Println("I got no idea what that means.")
		}
	}
}

func cthuluRoom() {
	fmt.Println("Here you see the great evil Cthulu.")
	fmt.Println("He, it, whatever stares at you and you go insane.")
	fmt.Println("Do you flee for your life or eat your head?")

	next := getInput("==> ")

	if strings.Contains(next, "flee") {
		start()
	} else if strings.Contains(next, "head") || strings.Contains(next, "eat") {
		dead("Well that was tasty!")
	} else {
		cthuluRoom()
	}
}

func dead(why string) {
	fmt.Println(why, "Good job!")
	os.Exit(0)
}

func start() {
	fmt.Println("You are in a dark room.")
	fmt.Println("There is a door to your right and left.")
	fmt.Println("Which one do you take?")

	next := getInput("==> ")

	switch next {
	case "left":
		bearRoom()
	case "right":
		cthuluRoom()
	default:
		dead("You stumble around the room until you starve.")
	}
}

func main() {
	start()
}
