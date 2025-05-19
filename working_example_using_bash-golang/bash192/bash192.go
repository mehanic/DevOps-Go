package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os/exec"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	MY_EMAIL             = "asweigart@gmail.com"
	BOT_EMAIL            = "imaptest@coffeeghost.net"
	BOT_EMAIL_PASSWORD    = "7|)6S1JS6>euu8p/nTlf"
	IMAP_SERVER          = "mail.coffeeghost.net:993"
	SMTP_SERVER          = "mail.coffeeghost.net:465"
	TORRENT_PROGRAM      = "C:\\Program Files (x86)\\qBittorrent\\qbittorrent.exe"
)

func main() {
	log.Println("Email bot started. Press Ctrl-C to quit.")

	for {
		instructions, err := getInstructionEmails()
		if err != nil {
			log.Println("Error getting instructions:", err)
			time.Sleep(15 * time.Minute)
			continue
		}

		for _, instruction := range instructions {
			log.Println("Doing instruction:", instruction)
			if err := parseInstructionEmail(instruction); err != nil {
				log.Println("Error parsing instruction:", err)
			}
		}

		log.Println("Done processing instructions. Pausing for 15 minutes.")
		time.Sleep(15 * time.Minute)
	}
}

func getInstructionEmails() ([]string, error) {
	// Connect to the IMAP server
	log.Println("Connecting to IMAP server...")
	imapClient, err := client.DialTLS(IMAP_SERVER, nil)
	if err != nil {
		return nil, err
	}
	defer imapClient.Logout()

	if err := imapClient.Login(BOT_EMAIL, BOT_EMAIL_PASSWORD); err != nil {
		return nil, err
	}
	imapClient.Select("INBOX", false)

	// Search for emails from MY_EMAIL
	criteria := imap.NewSearchCriteria()
	criteria.Header.Add("From", MY_EMAIL)

	UIDs, err := imapClient.UidSearch(criteria)
	if err != nil {
		return nil, err
	}

	var instructions []string
	for _, uid := range UIDs {
		messages := make(chan *imap.Message, 1)
		go func() {
			if err := imapClient.UidFetch([]uint32{uid}, []imap.FetchItem{imap.FetchEnvelope}, messages); err != nil {
				log.Println("Error fetching message:", err)
			}
		}()

		msg := <-messages
		if msg != nil {
			body := getEmailBody(imapClient, uid)
			instructions = append(instructions, body)
		}
	}

	// Delete the instruction emails if there are any
	if len(UIDs) > 0 {
		if err := imapClient.UidDelete(UIDs); err != nil {
			log.Println("Error deleting messages:", err)
		}
	}

	return instructions, nil
}

func getEmailBody(imapClient *client.Client, uid uint32) string {
	messages := make(chan *imap.Message, 1)
	go func() {
		if err := imapClient.UidFetch([]uint32{uid}, []imap.FetchItem{imap.FetchBody}, messages); err != nil {
			log.Println("Error fetching message body:", err)
		}
	}()

	msg := <-messages
	if msg != nil {
		if msg.Body != nil {
			// Read the email body
			mr := mail.NewReader(msg.Body)
			if _, err := mr.NextPart(); err == nil {
				if textPart, err := mr.NextPart(); err == nil {
					if body, err := textPart.Body(); err == nil {
						return string(body)
					}
				}
			}
		}
	}
	return ""
}

func parseInstructionEmail(instruction string) error {
	responseBody := "Subject: Instruction completed.\nInstruction received and completed.\nResponse:\n"

	// Check for magnet links
	if len(instruction) > 0 && instruction[:8] == "magnet:?" {
		cmd := exec.Command(TORRENT_PROGRAM, instruction)
		if err := cmd.Start(); err != nil {
			return err
		}
		responseBody += "Downloading magnet link.\n"
	}

	// Send email confirmation
	if err := sendEmail(responseBody); err != nil {
		return err
	}
	return nil
}

func sendEmail(body string) error {
	// Connect to the SMTP server
	log.Println("Connecting to SMTP server...")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "mail.coffeeghost.net",
	}
	conn, err := tls.Dial("tcp", SMTP_SERVER, tlsConfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, "mail.coffeeghost.net")
	if err != nil {
		return err
	}

	if err := c.Auth(smtp.PlainAuth("", BOT_EMAIL, BOT_EMAIL_PASSWORD, "mail.coffeeghost.net")); err != nil {
		return err
	}

	if err := c.Mail(BOT_EMAIL); err != nil {
		return err
	}

	if err := c.Rcpt(MY_EMAIL); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err := w.Write([]byte(body)); err != nil {
		return err
	}

	return c.Quit()
}

