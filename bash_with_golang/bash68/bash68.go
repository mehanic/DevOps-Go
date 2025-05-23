package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var pages = make(map[string]bool) // To keep track of visited pages

// getInternalLinks retrieves a list of all internal links found on a page
func getInternalLinks(bs *goquery.Document, includeUrl string) []string {
	internalLinks := []string{}
	parsedUrl, _ := url.Parse(includeUrl)

	// Finds all links that begin with a "/" or the same domain
	bs.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if !exists {
			return
		}
		if href != "" {
			if matched, _ := regexp.MatchString("^(/|.*"+parsedUrl.Host+")", href); matched {
				if href[0] == '/' {
					internalLinks = append(internalLinks, parsedUrl.Scheme+"://"+parsedUrl.Host+href)
				} else {
					internalLinks = append(internalLinks, href)
				}
			}
		}
	})

	return internalLinks
}

// getExternalLinks retrieves a list of all external links found on a page
func getExternalLinks(bs *goquery.Document, excludeUrl string) []string {
	externalLinks := []string{}
	parsedUrl, _ := url.Parse(excludeUrl)

	// Finds all links that start with "http" or "www" that do not contain the current URL
	bs.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if !exists {
			return
		}
		if href != "" {
			if matched, _ := regexp.MatchString("^(http|www)((?!"+parsedUrl.Host+").)*$", href); matched {
				externalLinks = append(externalLinks, href)
			}
		}
	})

	return externalLinks
}

// getRandomExternalLink retrieves a random external link from the specified page
func getRandomExternalLink(startingPage string) (string, error) {
	resp, err := http.Get(startingPage)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch URL: status code %d", resp.StatusCode)
	}

	bs, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	externalLinks := getExternalLinks(bs, startingPage)
	if len(externalLinks) == 0 {
		fmt.Println("No external links, looking around the site for one")
		parsedURL, err := url.Parse(startingPage)
		if err != nil {
			return "", err
		}
		domain := parsedURL.Scheme + "://" + parsedURL.Host
		internalLinks := getInternalLinks(bs, domain)
		if len(internalLinks) == 0 {
			return "", fmt.Errorf("no internal links found")
		}
		randomIndex := rand.Intn(len(internalLinks))
		return getRandomExternalLink(internalLinks[randomIndex])
	}

	randomIndex := rand.Intn(len(externalLinks))
	return externalLinks[randomIndex], nil
}

// followExternalOnly continuously follows external links starting from the initial site
func followExternalOnly(startingSite string) {
	externalLink, err := getRandomExternalLink(startingSite)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random external link is:", externalLink)
	followExternalOnly(externalLink)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	followExternalOnly("https://example.com")
}
