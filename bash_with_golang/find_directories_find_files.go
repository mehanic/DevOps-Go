package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
//	"strings"
)

// Define constants for paths and files
const (
	DIRB_PATH    = "/usr/bin/dirb"
	WORD_PATH    = "/usr/share/wordlists/dirb"
	WORD_LIST    = WORD_PATH + "/small.txt"
	EXT_LIST     = WORD_PATH + "/extensions_common.txt"
	RESOURCES    = "dirb_resources.txt"
)

// Regular expressions to parse the dirb output
var (
	dirRe   = regexp.MustCompile(`==> DIRECTORY: (.*?)`)
	reloRe  = regexp.MustCompile(`\+ (.*?) \(CODE: 3`)
	fileRe  = regexp.MustCompile(`\+ (.*?) \(CODE: [^3]`)
	fatalRe = regexp.MustCompile(`FATAL: (.*?)`)
)

func runCommand(cmd string) ([]byte, error) {
	// Run the command specified
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Check for fatal errors in the output
	if fatalRe.Match(out) {
		return nil, fmt.Errorf("fatal error: %s", fatalRe.FindStringSubmatch(string(out))[1])
	}

	return out, nil
}

func findDirectories(url string) ([]string, []string, error) {
	// Launch the dirb command against the specified url to find directories
	fmt.Printf("[*] Searching for directories at %s.\n", url)
	cmd := fmt.Sprintf("%s %s %s -Slrsiw", DIRB_PATH, url, WORD_LIST)
	resp, err := runCommand(cmd)
	if err != nil {
		return nil, nil, err
	}

	dirs := dirRe.FindAllStringSubmatch(string(resp), -1)
	files := fileRe.FindAllStringSubmatch(string(resp), -1)
	files = append(files, reloRe.FindAllStringSubmatch(string(resp), -1)...)

	var dirList []string
	for _, d := range dirs {
		dirList = append(dirList, d[1]+"/")
	}

	for _, dir := range dirList {
		fmt.Printf("[+] %s\n", dir)
	}

	return dirList, flatten(files), nil
}

func findFiles(url string) ([]string, error) {
	// Launch the dirb command against the specified URL to find files
	fmt.Printf("[*] Searching for files in %s.\n", url)
	cmd := fmt.Sprintf("%s %s %s -x %s -Srsiw", DIRB_PATH, url, WORD_LIST, EXT_LIST)
	resp, err := runCommand(cmd)
	if err != nil {
		return nil, err
	}

	files := fileRe.FindAllStringSubmatch(string(resp), -1)
	files = append(files, reloRe.FindAllStringSubmatch(string(resp), -1)...)

	for _, f := range files {
		fmt.Printf("[+] %s\n", f[1])
	}

	return flatten(files), nil
}

func flatten(matches [][]string) []string {
	var result []string
	for _, match := range matches {
		if len(match) > 1 {
			result = append(result, match[1])
		}
	}
	return result
}

func enumerate(target string) ([]string, error) {
	// Use dirb to enumerate content on the target
	target = target + "/"
	toSearch := []string{target}
	var directories []string
	var resources []string

	fmt.Println("\n[*] Recursively searching for directories.")
	for len(toSearch) != 0 {
		baseURL := toSearch[0]
		toSearch = toSearch[1:]

		dirs, files, err := findDirectories(baseURL)
		if err != nil {
			return nil, err
		}
		toSearch = append(toSearch, dirs...)
		directories = append(directories, baseURL)
		resources = append(resources, baseURL)
		resources = append(resources, files...)
	}

	fmt.Println("\n[*] Searching for files.")
	for _, url := range directories {
		files, err := findFiles(url)
		if err != nil {
			return nil, err
		}
		resources = append(resources, files...)
	}

	return resources, nil
}

func save(resources []string) error {
	// Save the resources to a file
	file, err := os.Create(RESOURCES)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, resource := range resources {
		_, err := writer.WriteString(resource + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go target_url")
		os.Exit(1)
	}

	resources, err := enumerate(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	err = save(resources)
	if err != nil {
		fmt.Printf("Error saving resources: %s\n", err)
		os.Exit(1)
	}
}

