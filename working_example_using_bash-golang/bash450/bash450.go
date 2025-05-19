package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

const (
	defaultNick    = "blb$$"
	defaultChannel = "testchannel"
	defaultServer  = "irc.freenode.net"
)

func main() {
	nick := defaultNick
	channel := defaultChannel
	server := defaultServer

	// Check for command line arguments to override default values
	if len(os.Args) > 1 {
		channel = os.Args[1]
	}
	if len(os.Args) > 2 {
		server = os.Args[2]
	}

	config := fmt.Sprintf("/tmp/irclog_%s", channel)

	// Write initial IRC commands to the config file
	if err := writeConfig(config, nick, channel); err != nil {
		fmt.Printf("Error writing config: %v\n", err)
		return
	}

	// Handle cleanup on exit
	cleanup := func() {
		os.Remove(config)
		fmt.Println("Cleanup done. Exiting.")
	}

	// Setup signal handling to gracefully cleanup
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cleanup()
		os.Exit(0)
	}()

	// Connect to the IRC server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:6667", server))
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	// Start listening for messages
	go readMessages(conn, config)

	// Send initial commands to the IRC server
	fmt.Fprintf(conn, "NICK %s\r\n", nick)
	fmt.Fprintf(conn, "USER %s +i * :%s\r\n", nick, os.Args[0])
	fmt.Fprintf(conn, "JOIN #%s\r\n", channel)

	// Keep the program running
	select {}
}

// writeConfig creates the config file with initial values
func writeConfig(configFile, nick, channel string) error {
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "NICK %s\nUSER %s +i * :%s\nJOIN #%s\n", nick, nick, os.Args[0], channel)
	return err
}

// readMessages listens for messages from the IRC server
func readMessages(conn net.Conn, config string) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		handleMessage(conn, message, config) // Pass conn to handleMessage
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
	}
}

// handleMessage processes incoming IRC messages
func handleMessage(conn net.Conn, message, config string) {
	switch {
	case message[:4] == "PING":
		fmt.Fprintf(conn, "PONG %s\r\n", message[5:])
		fmt.Println("Sent PONG response")

	case regexp.MustCompile(`^:(.*) PRIVMSG (.*) :(.*)$`).MatchString(message):
		// Handle PRIVMSG
		re := regexp.MustCompile(`^:(.*)!.* PRIVMSG (.*) :(.*)$`)
		match := re.FindStringSubmatch(message)
		if len(match) == 4 {
			sender := match[1]
			text := match[3]
			fmt.Printf("[%s] %s> %s\n", time.Now().Format("15:04"), sender, text)
		}

	default:
		fmt.Println(message)
	}
}

