package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define mail settings based on hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error retrieving hostname: %v\n", err)
		return
	}

	var mailer string
	switch {
	case strings.HasSuffix(hostname, ".company.com"):
		mailer = "mail" // Linux and BSD
	case strings.HasPrefix(hostname, "host1."):
		mailer = "mailx" // Solaris, BSD and some Linux
	case strings.HasPrefix(hostname, "host2."):
		mailer = "mailto" // Handy, if installed
	default:
		fmt.Println("Unknown mailer for this host.")
		return
	}

	recipients := "recipient1@example.com,recipient2@example.com"
	subject := fmt.Sprintf("Data from %s", filepath.Base(os.Args[0]))
	emailBody := "This is the body of the email.\n\nPlease find the attachment."

	// Specify the attachment file path
	attachment := "path/to/your/attachment.txt" // Change this to your actual attachment file path

	// Send the email
	if err := sendEmail(recipients, subject, emailBody, attachment); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	}
}

func sendEmail(recipients, subject, body, attachment string) error {
	// Set up the SMTP server configuration
	smtpHost := "smtp.example.com" // Change this to your SMTP server
	smtpPort := "587"               // Change this to your SMTP server port
	auth := smtp.PlainAuth("", "your_email@example.com", "your_password", smtpHost)

	// Create a buffer to hold the email content
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Create the email headers
	header := make(map[string]string)
	header["From"] = "your_email@example.com"
	header["To"] = recipients
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("multipart/mixed; boundary=%s", writer.Boundary())

	for key, value := range header {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	buf.WriteString("\r\n") // End of headers

	// Add the email body
	buf.WriteString(body + "\r\n")

	// Add the attachment
	if attachment != "" {
		part, err := writer.CreateFormFile("attachment", filepath.Base(attachment))
		if err != nil {
			return fmt.Errorf("failed to create form file for attachment: %v", err)
		}
		fileData, err := os.ReadFile(attachment)
		if err != nil {
			return fmt.Errorf("failed to read attachment file: %v", err)
		}
		part.Write(fileData)
	}

	// Close the writer to finalize the multipart message
	writer.Close()

	// Send the email
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, header["From"], strings.Split(recipients, ","), buf.Bytes()); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully.")
	return nil
}

