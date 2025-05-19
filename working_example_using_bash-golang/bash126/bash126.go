package main

import (
    "flag"
    "fmt"
    "io"
//    "net"
    "os"
    "os/exec"
    "time"
)

// Function to check if the host is reachable
func isHostUp(host string) bool {
    // Using ping command to check if the host is reachable
    cmd := exec.Command("ping", "-c", "1", host)
    err := cmd.Run()
    return err == nil
}

// Function to send an email
func sendEmail(email, host string) {
    subject := fmt.Sprintf("%s is up.", host)
    message := fmt.Sprintf("%s is up! %s\n", host, time.Now().Format(time.RFC1123))

    // Create a command to send email using mail command
    cmd := exec.Command("mailx", "-s", subject, email)

    // Create a pipe to send the message
    stdin, err := cmd.StdinPipe()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create stdin pipe: %v\n", err)
        return
    }

    // Start the mailx command
    if err := cmd.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to start mail command: %v\n", err)
        return
    }

    // Write the message to the stdin pipe and close it
    if _, err := io.WriteString(stdin, message); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to write to stdin: %v\n", err)
        return
    }
    stdin.Close() // Close the stdin pipe to signal end of input

    // Wait for the mail command to finish
    if err := cmd.Wait(); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to send email: %v\n", err)
    }
}

func usage() {
    fmt.Fprintln(os.Stderr, "Usage: email-when-up -e EMAIL_ADDRESS -h HOST")
    os.Exit(1)
}

func main() {
    var email string
    var host string

    // Parse command-line arguments
    flag.StringVar(&email, "e", "", "Email address to notify")
    flag.StringVar(&host, "h", "", "Host to check")
    flag.Parse()

    // Validate input
    if email == "" || host == "" {
        usage()
    }

    // Loop until the server is up
    const secondsToSleep = 60
    for {
        if isHostUp(host) {
            fmt.Printf("%s is up! %s\n", host, time.Now().Format(time.RFC1123))
            sendEmail(email, host)
            break
        } else {
            fmt.Printf("%s is still down. %s\n", host, time.Now().Format(time.RFC1123))
        }
        time.Sleep(secondsToSleep * time.Second)
    }
}

