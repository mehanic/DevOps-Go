package main

import (
	"fmt"
	"log"
	"calendar/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("My Birthday Party")
	if err != nil {
		log.Println(err)
		return
	}

	event.SetYear(2023)
	event.SetMonth(5)
	event.SetDay(15)

	fmt.Printf("Event: %s on %d-%02d-%02d\n", event.Title(), event.Year(), event.Month(), event.Day())
}

// go run main.go                                                                                                                                                ──(Fri,Jan17)─┘
// Event: My Birthday Party on 2023-05-15
