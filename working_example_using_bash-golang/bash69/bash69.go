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

var allExtLinks = make(map[string]bool) // Collects all external URLs found on the site
var allIntLinks = make(map[string]bool)  // Collects all internal URLs found on the site

// getInternalLinks retrieves a list of all internal links found on a page
func getInternalLinks(bs *goquery.Document, includeUrl string) []string {
	var internalLinks []string
	parsedUrl, _ := url.Parse(includeUrl)

	bs.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists && href != "" {
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
	var externalLinks []string
	parsedUrl, _ := url.Parse(excludeUrl)

	bs.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists && href != "" {
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
			return "", fmt.Errorf("failed to parse URL: %v", err)
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

// getAllExternalLinks collects a list of all external URLs found on the site
func getAllExternalLinks(siteUrl string) error {
	resp, err := http.Get(siteUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	parsedURL, err := url.Parse(siteUrl)
	if err != nil {
		return err
	}
	domain := parsedURL.Scheme + "://" + parsedURL.Host

	bs, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	internalLinks := getInternalLinks(bs, domain)
	externalLinks := getExternalLinks(bs, domain)

	for _, link := range externalLinks {
		if !allExtLinks[link] {
			allExtLinks[link] = true
			fmt.Println(link)
		}
	}
	for _, link := range internalLinks {
		if !allIntLinks[link] {
			allIntLinks[link] = true
			if err := getAllExternalLinks(link); err != nil {
				return err
			}
		}
	}
	return nil
}
func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	startingSite := "http://oreilly.com"
	followExternalOnly(startingSite)

	allIntLinks[startingSite] = true // To avoid re-fetching the starting site
	if err := getAllExternalLinks(startingSite); err != nil {
		log.Fatal(err)
	}
}

