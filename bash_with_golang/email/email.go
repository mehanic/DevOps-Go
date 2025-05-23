package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Sender and recipient information
	from := "ryan@pythonscraping.com"
	to := "webmaster@pythonscraping.com"
	subject := "An Email Alert"
	body := "The body of the email is here"

	// Compose the email message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Set up the SMTP server configuration
	smtpHost := "localhost"
	smtpPort := "25" // Port 25 for local SMTP server

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	fmt.Println("Email sent successfully")
}

