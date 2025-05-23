package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/miekg/dns"
)

// getMXRecord retrieves the MX record for the given domain
func getMXRecord(domain string) (string, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return "", err
	}
	return mxRecords[0].Host, nil
}

// sendMail sends an email with an optional attachment
func sendMail(sendFrom, sendTo, subject, text, filename string) error {
	// Prepare the email headers
	msg := &bytes.Buffer{}
	msg.WriteString("From: " + sendFrom + "\n")
	msg.WriteString("To: " + sendTo + "\n")
	msg.WriteString("Subject: " + subject + "\n")
	msg.WriteString("MIME-Version: 1.0\n")
	msg.WriteString("Content-Type: multipart/mixed; boundary=\"boundary\"\n\n")
	msg.WriteString("--boundary\n")
	msg.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\n\n")
	msg.WriteString(text + "\n\n")
	
	if filename != "" {
		// Attach the file
		fileData, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		msg.WriteString("--boundary\n")
		msg.WriteString("Content-Type: application/octet-stream\n")
		msg.WriteString("Content-Disposition: attachment; filename=\"" + filepath.Base(filename) + "\"\n")
		msg.WriteString("\n")
		msg.Write(fileData)
		msg.WriteString("\n\n")
		msg.WriteString("--boundary--\n")
	}

	// Get the recipient domain and MX record
	domain := sendTo[strings.Index(sendTo, "@")+1:]
	server, err := getMXRecord(domain)
	if err != nil {
		return err
	}

	// Send the email
	smtpServer := smtp.ServerInfo{
		Addr: server,
	}
	auth := smtp.PlainAuth("", sendFrom, "", smtpServer.Addr)

	if err := smtp.SendMail(smtpServer.Addr, auth, sendFrom, []string{sendTo}, msg.Bytes()); err != nil {
		return err
	}

	return nil
}

// usage displays how to use the program
func usage() {
	fmt.Println("Usage: send_email from to subject text [attachment]")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 5 {
		usage()
	}

	sendFrom := os.Args[1]
	sendTo := os.Args[2]
	subject := os.Args[3]
	text := os.Args[4]
	var filename string

	if len(os.Args) == 6 {
		filename = os.Args[5]
	}

	if err := sendMail(sendFrom, sendTo, subject, text, filename); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}

