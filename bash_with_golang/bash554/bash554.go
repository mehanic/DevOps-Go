package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "sort"
    "strings"
)

// downloadLinks fetches the content from the URL and extracts unique domain names
func downloadLinks(url string) ([]string, error) {
    // Define the regex for matching URLs
    regex := `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]*[-A-Za-z0-9+&@#/%=~_|]`

    // Fetch the URL content
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching URL %s: %v", url, err)
    }
    defer resp.Body.Close()

    // Read the body content
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading body: %v", err)
    }

    // Compile the regex
    re := regexp.MustCompile(regex)
    matches := re.FindAllString(string(body), -1)

    // Use a map to keep track of unique domains
    domains := make(map[string]struct{})

    // Extract domains from the matches
    for _, match := range matches {
        parts := strings.Split(match, "/")
        if len(parts) > 2 {
            domain := parts[2] // Extract the domain part
            domains[domain] = struct{}{} // Add to map
        }
    }

    // Convert map keys to a slice for sorting
    uniqueDomains := make([]string, 0, len(domains))
    for domain := range domains {
        uniqueDomains = append(uniqueDomains, domain)
    }

    // Sort the unique domains
    sort.Strings(uniqueDomains)
    return uniqueDomains, nil
}

func main() {
    urls := []string{
        "http://www.cisco.com/",
        "https://www.yahoo.com/",
    }

    for _, url := range urls {
        uniqueDomains, err := downloadLinks(url)
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("Unique domains from %s:\n", url)
        for _, domain := range uniqueDomains {
            fmt.Println(domain)
        }
    }
}

