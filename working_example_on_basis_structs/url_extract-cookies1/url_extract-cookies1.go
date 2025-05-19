package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
//	"regexp"
	"strings"
	"time"
)

//-----------------------------------------------------------------------------
var (
	client *http.Client
)

//-----------------------------------------------------------------------------
func loadFile(filename string, insensitive bool) map[string]struct{} {
	// Load each line of a file into a set to get only unique items. Case-insensitive if required.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("[-] Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	items := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "/")
		if insensitive {
			line = strings.ToLower(line)
		}
		items[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[-] Error reading file: %v\n", err)
	}

	fmt.Printf("[+] Loaded %d items.\n", len(items))
	return items
}

//-----------------------------------------------------------------------------
func buildFilenames(nameList, extList map[string]struct{}) map[string]struct{} {
	// Build filenames from the name list and the extension list.
	names := make(map[string]struct{})
	if len(extList) == 0 {
		names = nameList
	} else {
		for name := range nameList {
			for ext := range extList {
				names[fmt.Sprintf("%s.%s", name, strings.TrimPrefix(ext, "."))] = struct{}{}
			}
		}
	}

	fmt.Printf("[+] Built %d filenames.\n", len(names))
	return names
}

//-----------------------------------------------------------------------------
func head(url string) bool {
	// Make a HEAD request to check if the URL exists (not returning 404).
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Printf("[-] Error creating request: %v\n", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 301 || resp.StatusCode == 302 {
		fmt.Printf("[+] Found %s --> %s (%d)\n", url, resp.Header.Get("Location"), resp.StatusCode)
		return true
	} else if resp.StatusCode != 404 {
		fmt.Printf("[+] Found %s (%d)\n", url, resp.StatusCode)
		return true
	}

	return false
}

//-----------------------------------------------------------------------------
func check(baseURL string, pathList map[string]struct{}, isDir bool) []string {
	// Check each URL in the path list to see if it exists.
	var resources []string
	baseURL = strings.TrimRight(baseURL, "/")

	for path := range pathList {
		if isDir {
			path = fmt.Sprintf("%s/", strings.TrimRight(path, "/"))
		}
		url := fmt.Sprintf("%s/%s", baseURL, path)
		if head(url) {
			resources = append(resources, url)
		}
	}

	return resources
}

//-----------------------------------------------------------------------------
func buildLists(directoryFile, nameFile, extensionFile string, insensitive bool) (map[string]struct{}, map[string]struct{}) {
	// Load directory, name, and extensions lists into memory.
	fmt.Println("[*] Getting enumeration lists.")

	fmt.Printf("[+] Loading directory file %s.\n", directoryFile)
	directoryList := loadFile(directoryFile, insensitive)

	var nameList, extList map[string]struct{}
	nameList = make(map[string]struct{})
	extList = make(map[string]struct{})

	if nameFile != "" {
		fmt.Printf("[+] Loading name file %s.\n", nameFile)
		nameList = loadFile(nameFile, insensitive)

		if extensionFile != "" {
			fmt.Printf("[+] Loading extension file %s.\n", extensionFile)
			extList = loadFile(extensionFile, insensitive)
		}
	}

	filenames := buildFilenames(nameList, extList)
	return directoryList, filenames
}

//-----------------------------------------------------------------------------
func enumerate(baseURL string, directoryList, filenames map[string]struct{}) []string {
	// Enumerate directories and files on the web server.
	fmt.Println("[*] Enumerating resources.")
	var toSearch, directories, resources []string
	toSearch = append(toSearch, baseURL)

	fmt.Println("[*] Recursively searching for directories.")
	for len(toSearch) > 0 {
		baseURL = toSearch[0]
		toSearch = toSearch[1:]

		fmt.Printf("[*] Searching for directories in %s\n", baseURL)
		foundDirs := check(baseURL, directoryList, true)
		toSearch = append(toSearch, foundDirs...)
		directories = append(directories, baseURL)
		resources = append(resources, baseURL)
	}

	if len(filenames) > 0 {
		fmt.Println("[*] Searching for files.")
		for _, dir := range directories {
			resources = append(resources, check(dir, filenames, false)...)
		}
	}

	return resources
}

//-----------------------------------------------------------------------------
func saveResources(filename string, resources []string) {
	// Save identified resources to a file.
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("[-] Error creating file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	for _, resource := range resources {
		file.WriteString(resource + "\n")
	}
}

//-----------------------------------------------------------------------------
func main() {
	// Define and parse flags
	baseURL := flag.String("url", "", "Target web server to enumerate")
	directoryFile := flag.String("dir", "", "List of directories to enumerate")
	nameFile := flag.String("n", "", "List of filenames to enumerate")
	extensionFile := flag.String("e", "", "List of extensions to append to filenames")
	caseInsensitive := flag.Bool("i", false, "Use case-insensitive matching")
	outputFile := flag.String("o", "resources.txt", "Output file to save the results")
	flag.Parse()

	if *baseURL == "" || *directoryFile == "" {
		fmt.Println("Usage: ./web_discover -url <base_url> -dir <directory_file>")
		return
	}

	// Set up the HTTP client
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Ensure the base URL is reachable
	if !head(*baseURL) {
		fmt.Printf("[-] Unable to connect to %s.\n", *baseURL)
		return
	}

	// Build directory list and filenames
	directoryList, filenames := buildLists(*directoryFile, *nameFile, *extensionFile, *caseInsensitive)

	// Enumerate resources
	resources := enumerate(*baseURL, directoryList, filenames)

	// Save the results
	saveResources(*outputFile, resources)
}

