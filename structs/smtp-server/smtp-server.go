package main

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

// GetMXRecord retrieves the MX record for the given domain.
func getMXRecord(domain string) (string, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return "", err
	}
	return mxRecords[0].Host, nil
}

// CustomSMTPServer defines a custom SMTP server.
type CustomSMTPServer struct {
	Addr string
}

// NewCustomSMTPServer creates a new SMTP server instance.
func NewCustomSMTPServer(addr string) *CustomSMTPServer {
	return &CustomSMTPServer{Addr: addr}
}

// Start begins listening for incoming SMTP connections.
func (server *CustomSMTPServer) Start() {
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", server.Addr, err)
	}
	defer listener.Close()

	log.Printf("[+] Server listening on %s", server.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go server.handleConnection(conn)
	}
}

// handleConnection processes incoming SMTP connections.
func (server *CustomSMTPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Here, you can handle the SMTP protocol and parse the incoming messages
	// For simplicity, we're just logging incoming connections
	log.Printf("New connection from %s", conn.RemoteAddr().String())

	// In a real implementation, you'd read from conn and process the email data.
	// For this example, we will just close the connection.
}

// ProcessMessage processes the incoming message and sends it to the appropriate recipients.
func processMessage(mailFrom string, rcptTo []string, data []byte) {
	for _, recipient := range rcptTo {
		fmt.Printf("[*] Sending message to %s.\n", recipient)

		domain := recipient[strings.Index(recipient, "@")+1:]
		mx, err := getMXRecord(domain)
		if err != nil {
			log.Printf("[-] Failed to get MX record for %s: %v", domain, err)
			continue
		}

		// Create SMTP client
		smtpClient, err := smtp.Dial(fmt.Sprintf("%s:25", mx))
		if err != nil {
			log.Printf("[-] Failed to connect to %s: %v", mx, err)
			continue
		}
		defer smtpClient.Close()

		if err := smtpClient.Mail(mailFrom); err != nil {
			log.Printf("[-] Failed to set mail from %s: %v", mailFrom, err)
			continue
		}
		if err := smtpClient.Rcpt(recipient); err != nil {
			log.Printf("[-] Failed to set recipient %s: %v", recipient, err)
			continue
		}

		// Send the email data
		writer, err := smtpClient.Data()
		if err != nil {
			log.Printf("[-] Failed to send data: %v", err)
			continue
		}

		if _, err := writer.Write(data); err != nil {
			log.Printf("[-] Failed to write data: %v", err)
		}

		if err := writer.Close(); err != nil {
			log.Printf("[-] Failed to close writer: %v", err)
		}
	}
}

func main() {
	const port = "2525"

	server := NewCustomSMTPServer("127.0.0.1:" + port)
	go server.Start()

	// Keep the server running indefinitely
	select {}
}

