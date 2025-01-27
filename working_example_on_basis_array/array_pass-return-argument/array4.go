package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}

	name := args[0]

	moods := [...]string{
		"happy 😀", "good 👍", "awesome 😎",
		"sad 😞", "bad 👎", "terrible 😩",
	}

	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[n])
}

// max feels good 👍
