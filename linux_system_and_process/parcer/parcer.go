package main

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"
)

type ColorizeParser struct {
	pieces     []string
	colorindex int
	needcolor  bool
	usefonts   bool
}

var fontDataMap = map[string][2]string{
	"comment":  {"<font color='green'><i>", "</i></font>"},
	"string":   {"<font color='olive'>", "</font>"},
	"keyword":  {"<font color='navy'><b>", "</b></font>"},
	"function": {"<font color='teal'><b>", "</b></font>"},
	"class":    {"<font color='blue'><b>", "</b></font>"},
}

// Initialize a new ColorizeParser
func NewColorizeParser(usefonts bool) *ColorizeParser {
	return &ColorizeParser{usefonts: usefonts}
}

func (p *ColorizeParser) reset() {
	p.pieces = []string{}
	p.colorindex = 0
	p.needcolor = false
}

func (p *ColorizeParser) HTMLfontify(text string) string {
	// Dummy tokenization logic, replace with pyfontify equivalent
	fontmap := []struct {
		token string
		start int
		end   int
	}{
		{"keyword", 0, 5},  // Example token
		{"string", 6, 10},  // Example token
	}

	for i := len(fontmap) - 1; i >= 0; i-- {
		token := fontmap[i].token
		start := fontmap[i].start
		end := fontmap[i].end

		if p.usefonts {
			text = text[:start] + fontDataMap[token][0] + text[start:end] + fontDataMap[token][1] + text[end:]
		} else {
			text = text[:start] + fmt.Sprintf("<span class='py%s'>%s</span>", token, text[start:end]) + text[end:]
		}
	}

	return text
}

func (p *ColorizeParser) flushcolor() {
	if p.colorindex > 0 {
		buffer := strings.Join(p.pieces[p.colorindex:], "")
		p.pieces = p.pieces[:p.colorindex]
		p.colorindex = 0
		p.pieces = append(p.pieces, p.HTMLfontify(buffer))
	}
}

// This is a placeholder for handling start tags in the HTML processor
func (p *ColorizeParser) unknownStartTag(tag string, attrs map[string]string) {
	p.flushcolor()
	// Logic to handle the start tag
	if p.needcolor {
		p.colorindex = len(p.pieces)
	}
}

// This is a placeholder for handling end tags in the HTML processor
func (p *ColorizeParser) unknownEndTag(tag string) {
	p.flushcolor()
	// Logic to handle the end tag
	if p.needcolor {
		p.colorindex = len(p.pieces)
	}
}

// Read and process the file
func process(filename string, usefonts bool) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parser := NewColorizeParser(usefonts)
	parser.reset()

	var output strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		parser.pieces = append(parser.pieces, html.EscapeString(line))
	}

	parser.flushcolor()
	return output.String(), nil
}

// Process a directory of files or a single file
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: colorize directory-or-file")
		return
	}

	filedir := os.Args[1]
	usefonts := len(os.Args) > 2 && os.Args[2] == "1"

	if stat, err := os.Stat(filedir); err == nil && stat.IsDir() {
		filepath.Walk(filedir, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".html") {
				fmt.Printf("Colorizing %s\n", filepath.Base(path))
				_, err := process(path, usefonts)
				if err != nil {
					fmt.Println("Error processing file:", err)
				}
			}
			return nil
		})
	} else {
		fmt.Printf("Colorizing %s\n", filepath.Base(filedir))
		_, err := process(filedir, usefonts)
		if err != nil {
			fmt.Println("Error processing file:", err)
		}
	}
}

