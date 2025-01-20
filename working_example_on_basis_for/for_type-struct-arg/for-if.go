package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the AnonymousSurvey struct
type AnonymousSurvey struct {
	question  string
	responses []string
}

// Method to show the survey question
func (s *AnonymousSurvey) showQuestion() {
	fmt.Println(s.question)
}

// Method to store a response
func (s *AnonymousSurvey) storeResponse(response string) {
	s.responses = append(s.responses, response)
}

// Method to show survey results
func (s *AnonymousSurvey) showResults() {
	fmt.Println("\nSurvey Results:")
	for _, response := range s.responses {
		fmt.Println("- " + response)
	}
}

func main() {
	// Define a question, and make a survey
	question := "What language did you first learn to speak?"
	mySurvey := AnonymousSurvey{question: question}

	// Show the question, and store responses
	mySurvey.showQuestion()
	fmt.Println("Enter 'q' at any time to quit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		// Read input from the user
		fmt.Print("Language: ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if response == "q" {
			break
		}

		// Store the response
		mySurvey.storeResponse(response)
	}

	// Show the survey results
	fmt.Println("\nThank you to everyone who participated in the survey!")
	mySurvey.showResults()
}

// $ go run /home/mehanic/structure/for/for-if/for-if.go
// What language did you first learn to speak?
// Enter 'q' at any time to quit.

// Language: chinese
// Language: uk
// Language: q

// Thank you to everyone who participated in the survey!

// Survey Results:
// - chinese
// - uk
