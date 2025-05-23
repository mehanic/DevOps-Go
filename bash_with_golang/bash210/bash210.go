package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Load the Excel file
	file, err := excelize.OpenFile("duesRecords.xlsx")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Get the sheet by name
	sheetName := "Sheet1"
	rows, err := file.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Error getting sheet: %v", err)
	}

	// Get the last column and latest month
	lastCol := len(rows[0])
	latestMonth := rows[0][lastCol-1]

	// Create a map for unpaid members
	unpaidMembers := make(map[string]string)

	// Check each member's payment status
	for i, row := range rows[1:] { // Skip the header row
		if len(row) < lastCol {
			continue
		}
		payment := row[lastCol-1]
		if payment != "paid" {
			name := row[0]
			email := row[1]
			unpaidMembers[name] = email
		}
	}

	// Get email password from command-line argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <email_password>", os.Args[0])
	}
	emailPassword := os.Args[1]

	// Set up the email server (Gmail in this case)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "my_email_address@gmail.com"
	auth := smtp.PlainAuth("", senderEmail, emailPassword, smtpHost)

	// Send out reminder emails to unpaid members
	for name, email := range unpaidMembers {
		subject := fmt.Sprintf("%s dues unpaid", latestMonth)
		body := fmt.Sprintf("Dear %s,\n\nRecords show that you have not paid dues for %s. Please make this payment as soon as possible. Thank you!", name, latestMonth)
		message := []byte("Subject: " + subject + "\r\n" +
			"\r\n" + body)

		fmt.Printf("Sending email to %s...\n", email)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{email}, message)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		} else {
			fmt.Printf("Email sent to %s successfully!\n", email)
		}
	}
}

