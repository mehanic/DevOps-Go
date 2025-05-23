package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Account information for Twilio
const accountSID = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
const authToken = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
const myNumber = "+15559998888"
const twilioNumber = "+15552225678"

// Function to send an SMS
func textMyself(message string) {
	// Create a new Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	// Create a message
	params := &openapi.CreateMessageParams{}
	params.SetTo(myNumber)
	params.SetFrom(twilioNumber)
	params.SetBody(message)

	// Send the message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	} else {
		fmt.Println("Message sent successfully, SID:", *resp.Sid)
	}
}

func main() {
	// Example usage of the textMyself function
	message := "Hello, this is a test message!"
	textMyself(message)
}

