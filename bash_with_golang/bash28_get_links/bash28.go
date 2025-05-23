package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Function to get all links from a Wikipedia page
func getLinks(articleUrl string) ([]string, error) {
	resp, err := http.Get("http://en.wikipedia.org" + articleUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	doc.Find("div#bodyContent a").Each(func(index int, item *goquery.Selection) {
		href, exists := item.Attr("href")
		if exists {
			// Using a regex to match valid Wikipedia URLs (without colons)
			match, _ := regexp.MatchString("^(/wiki/)((?!:).)*$", href)
			if match {
				links = append(links, href)
			}
		}
	})
	return links, nil
}

// Function to get history IPs from Wikipedia page
func getHistoryIPs(pageUrl string) (map[string]bool, error) {
	pageUrl = strings.Replace(pageUrl, "/wiki/", "", 1)
	historyUrl := "http://en.wikipedia.org/w/index.php?title=" + pageUrl + "&action=history"
	fmt.Println("History URL is: " + historyUrl)

	resp, err := http.Get(historyUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	ipAddresses := make(map[string]bool)
	doc.Find("a.mw-anonuserlink").Each(func(index int, item *goquery.Selection) {
		ip := item.Text()
		ipAddresses[ip] = true
	})

	return ipAddresses, nil
}

func getCountry(ipAddress string) (string, error) {
	resp, err := http.Get("http://ip-api.com/json/" + ipAddress)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	country, ok := result["countryCode"].(string)  // Adjusted for ip-api.com
	if !ok {
		return "", fmt.Errorf("could not retrieve country code for IP %s", ipAddress)
	}
	return country, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	links, err := getLinks("/wiki/Python_(programming_language)")
	if err != nil {
		log.Fatal(err)
	}

	// Print the initial links
	fmt.Println("Initial Links Found:", links)

	for len(links) > 0 {
		for _, link := range links {
			fmt.Println("-------------------")
			historyIPs, err := getHistoryIPs(link)
			if err != nil {
				log.Println("Error getting history IPs:", err)
				continue
			}

			// Print history IPs found
			fmt.Println("History IPs Found:", historyIPs)

			for historyIP := range historyIPs {
				country, err := getCountry(historyIP)
				if err != nil {
					log.Println("Error getting country:", err)
				} else {
					fmt.Printf("%s is from %s\n", historyIP, country)
				}
			}
		}

		// Select a random new link to crawl next
		newLink := links[rand.Intn(len(links))]
		links, err = getLinks(newLink)
		if err != nil {
			log.Fatal(err)
		}
	}
}

