package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	downloadDirectory = "downloaded"
	baseUrl           = "http://pythonscraping.com"
)

func main() {
	// Fetch the HTML page
	resp, err := http.Get("http://www.pythonscraping.com")
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	// Find all elements with src attributes
	doc.Find("[src]").Each(func(index int, element *goquery.Selection) {
		src, exists := element.Attr("src")
		if exists {
			fileUrl := getAbsoluteURL(baseUrl, src)
			if fileUrl != "" {
				fmt.Println(fileUrl)
				downloadPath := getDownloadPath(baseUrl, fileUrl, downloadDirectory)
				err := downloadFile(fileUrl, downloadPath)
				if err != nil {
					log.Printf("Failed to download file: %v", err)
				}
			}
		}
	})
}

// getAbsoluteURL constructs the absolute URL from the source
func getAbsoluteURL(baseUrl, source string) string {
	if strings.HasPrefix(source, "http://www.") {
		return "http://" + source[11:]
	} else if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return source
	} else if strings.HasPrefix(source, "www.") {
		return "http://" + source[4:]
	}
	url := baseUrl + "/" + source
	if !strings.HasPrefix(url, baseUrl) {
		return ""
	}
	return url
}

// getDownloadPath generates the local path for the downloaded file
func getDownloadPath(baseUrl, absoluteUrl, downloadDirectory string) string {
	path := strings.Replace(absoluteUrl, "www.", "", 1)
	path = strings.Replace(path, baseUrl, "", 1)
	path = filepath.Join(downloadDirectory, path)

	// Create the directory if it does not exist
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return path
}

// downloadFile downloads the file from the given URL to the specified path
func downloadFile(url, filePath string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the data to the file
	_, err = io.Copy(out, resp.Body)
	return err
}

