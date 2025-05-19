package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url" // Import the url package
	"os"
	"regexp"
	"sync"
)

var (
	cookieRe   = regexp.MustCompile(`.*?=([0-9A-Za-z%]+)--([0-9a-f]+);`)
	workerCount = 4
)

// unquote decodes the URL-encoded data
func unquote(data string) string {
	udata, err := url.QueryUnescape(data)
	if err != nil {
		return data
	}
	if data == udata {
		return udata
	}
	return unquote(udata)
}

// extractSessionDigest extracts the session and digest from the cookie
func extractSessionDigest(cookie string) (string, string) {
	matches := cookieRe.FindStringSubmatch(cookie)
	if matches != nil {
		return unquote(matches[1]), matches[2]
	}
	return "", ""
}

// worker processes URLs from the urlQueue and extracts session cookies
func worker(urlQueue <-chan string, cookieChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[*] Starting new worker thread.")

	for url := range urlQueue {
		fmt.Printf("[*] Accessing %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("[-] Could not access %s\n", url)
			continue
		}
		defer resp.Body.Close()

		cookie := resp.Header.Get("Set-Cookie")
		if cookie != "" {
			session, digest := extractSessionDigest(cookie)
			if session != "" {
				fmt.Printf("[*] Found matching cookie for %s.\n", url)
				cookieChan <- fmt.Sprintf("%s::%s::%s", url, session, digest)
			}
		}
	}
}

// processFile reads URLs from a file and processes them
func processFile(filename string) {
	urlQueue := make(chan string, 100)
	cookieChan := make(chan string)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(urlQueue, cookieChan, &wg)
	}

	// Load the URLs into the queue
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[-] Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := "http://" + scanner.Text()
		urlQueue <- url
	}
	close(urlQueue)

	// Wait for all workers to finish
	wg.Wait()
	close(cookieChan)

	// Write the cookies to a file
	outFilename := "results_" + filename
	outfile, err := os.Create(outFilename)
	if err != nil {
		fmt.Println("[-] Error creating output file:", err)
		return
	}
	defer outfile.Close()

	for cookie := range cookieChan {
		_, err := outfile.WriteString(cookie + "\n")
		if err != nil {
			fmt.Println("[-] Error writing to output file:", err)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: stage_find <url_file>")
		return
	}

	filename := os.Args[1]
	fmt.Printf("[*] Processing file %s\n", filename)
	processFile(filename)
}

