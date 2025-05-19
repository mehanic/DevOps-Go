package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	nick := "blb$$"
	channel := "testchannel"
	server := "irc.freenode.net"

	// Check for command-line arguments
	if len(os.Args) > 1 {
		channel = os.Args[1]
	}
	if len(os.Args) > 2 {
		server = os.Args[2]
	}

	configFile := fmt.Sprintf("/tmp/irclog_%s", channel)
	config, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Error creating config file:", err)
		return
	}
	defer os.Remove(configFile) // Remove the file on exit
	defer config.Close()

	// Write NICK and USER commands to the config file
	fmt.Fprintf(config, "NICK %s\n", nick)
	fmt.Fprintf(config, "USER %s +i * :%s\n", nick, os.Args[0])
	fmt.Fprintf(config, "JOIN #%s\n", channel)

	// Set up a TCP connection to the IRC server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:6667", server))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Set up signal handling to clean up on exit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		os.Remove(configFile) // Clean up config file on exit
		os.Exit(0)
	}()

	// Start reading from the server
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Send the initial commands
	fmt.Fprintf(writer, "NICK %s\r\n", nick)
	fmt.Fprintf(writer, "USER %s +i * :%s\r\n", nick, os.Args[0])
	fmt.Fprintf(writer, "JOIN #%s\r\n", channel)
	writer.Flush()

	// Listen for incoming messages
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			break
		}
		message = strings.TrimSpace(message)
		fmt.Println(message) // Print raw message to console

		// Handle PING messages
		if strings.HasPrefix(message, "PING") {
			pongResponse := fmt.Sprintf("PONG %s\r\n", message[5:])
			writer.WriteString(pongResponse)
			writer.Flush()
			continue
		}

		// Handle PRIVMSG and log it to the config file
		if strings.Contains(message, "PRIVMSG") {
			parts := strings.SplitN(message, "PRIVMSG", 2)
			if len(parts) == 2 {
				user := strings.TrimPrefix(parts[0], ":")
				user = strings.Split(user, "!")[0] // Extract the nickname
				msg := strings.TrimSpace(parts[1])
				if strings.HasPrefix(msg, ":") {
					msg = strings.TrimPrefix(msg, ":")
				}
				// Format message with current time
				logMessage := fmt.Sprintf("[%s] %s> %s\n", time.Now().Format("15:04"), user, msg)
				fmt.Fprint(config, logMessage) // Log to file
				fmt.Print(logMessage)           // Print to console
			}
		}
	}
}

