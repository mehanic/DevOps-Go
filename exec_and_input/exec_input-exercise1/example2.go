package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Greeting запрашивает его имя пользователя и номер стола.
func Greeting() error {
	questions := []string{"What is your name?", "What table do you need?"}
	var answers []string

	scanner := bufio.NewScanner(os.Stdin)
	for _, question := range questions {
		fmt.Println(question)

		scanner.Scan()
		text := scanner.Text()

		if len(text) == 0 {
			return errors.New("input cannot be empty")
		}

		answers = append(answers, strings.TrimSpace(text))
	}

	if scanner.Err() != nil {
		return fmt.Errorf("error while scanning input: %v", scanner.Err())
	}

	table, err := strconv.Atoi(answers[1])
	if err != nil {
		return fmt.Errorf("error converting table number to integer: %v", err)
	}

	message := fmt.Sprintf("Hello, %s! Your table is %d.", answers[0], table)
	fmt.Println(message)

	return nil
}

// main function: entry point of the program
func main() {
	if err := Greeting(); err != nil {
		fmt.Println("Error:", err)
	}
}


// What is your name?
// max
// What table do you need?
// 6
// Hello, max! Your table is 6.
