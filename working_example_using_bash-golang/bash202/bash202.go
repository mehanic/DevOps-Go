package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var capitals = map[string]string{
	"Alabama":       "Montgomery",
	"Alaska":        "Juneau",
	"Arizona":       "Phoenix",
	"Arkansas":      "Little Rock",
	"California":    "Sacramento",
	"Colorado":      "Denver",
	"Connecticut":   "Hartford",
	"Delaware":      "Dover",
	"Florida":       "Tallahassee",
	"Georgia":       "Atlanta",
	"Hawaii":        "Honolulu",
	"Idaho":         "Boise",
	"Illinois":      "Springfield",
	"Indiana":       "Indianapolis",
	"Iowa":          "Des Moines",
	"Kansas":        "Topeka",
	"Kentucky":      "Frankfort",
	"Louisiana":     "Baton Rouge",
	"Maine":         "Augusta",
	"Maryland":      "Annapolis",
	"Massachusetts": "Boston",
	"Michigan":      "Lansing",
	"Minnesota":     "Saint Paul",
	// Add more state-capital pairs here...
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for quizNum := 1; quizNum <= 35; quizNum++ {
		quizFileName := fmt.Sprintf("capitalsquiz%d.txt", quizNum)
		answerKeyFileName := fmt.Sprintf("capitalsquiz_answers%d.txt", quizNum)

		quizFile, err := os.Create(quizFileName)
		if err != nil {
			fmt.Println("Error creating quiz file:", err)
			continue
		}
		defer quizFile.Close()

		answerKeyFile, err := os.Create(answerKeyFileName)
		if err != nil {
			fmt.Println("Error creating answer key file:", err)
			continue
		}
		defer answerKeyFile.Close()

		quizHeader := "Name:\n\nDate:\n\nPeriod:\n\n" + strings.Repeat(" ", 20) + fmt.Sprintf("State Capitals Quiz (Form %d)\n\n", quizNum)
		quizFile.WriteString(quizHeader)

		// Shuffle states
		states := make([]string, 0, len(capitals))
		for state := range capitals {
			states = append(states, state)
		}
		rand.Shuffle(len(states), func(i, j int) {
			states[i], states[j] = states[j], states[i]
		})

		// Create quiz questions
		for questionNum := 1; questionNum <= 50; questionNum++ {
			state := states[questionNum-1]
			correctAnswer := capitals[state]

			// Get three random incorrect answers
			wrongAnswers := make([]string, 0, len(capitals)-1)
			for _, capital := range capitals {
				if capital != correctAnswer {
					wrongAnswers = append(wrongAnswers, capital)
				}
			}
			rand.Shuffle(len(wrongAnswers), func(i, j int) {
				wrongAnswers[i], wrongAnswers[j] = wrongAnswers[j], wrongAnswers[i]
			})
			wrongAnswers = wrongAnswers[:3]

			// Combine correct and wrong answers, and shuffle them
			answerOptions := append(wrongAnswers, correctAnswer)
			rand.Shuffle(len(answerOptions), func(i, j int) {
				answerOptions[i], answerOptions[j] = answerOptions[j], answerOptions[i]
			})

			// Write question to quiz file
			quizFile.WriteString(fmt.Sprintf("%d. What is the capital of %s?\n", questionNum, state))
			for i, option := range answerOptions {
				quizFile.WriteString(fmt.Sprintf("    %s. %s\n", string('A'+i), option))
			}
			quizFile.WriteString("\n")

			// Write answer to answer key file
			for i, option := range answerOptions {
				if option == correctAnswer {
					answerKeyFile.WriteString(fmt.Sprintf("%d. %s\n", questionNum, string('A'+i)))
					break
				}
			}
		}
	}
}

