package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Dialectizer - Base struct for processing dialects
type Dialectizer struct {
	subs         [][2]string
	output       strings.Builder
	verbatimMode int
}

// Process the text using regular expressions
func (d *Dialectizer) process(text string) string {
	for _, sub := range d.subs {
		re := regexp.MustCompile(sub[0])
		text = re.ReplaceAllString(text, sub[1])
	}
	return text
}

// Handle HTML <pre> tag (start)
func (d *Dialectizer) startPre(attrs string) {
	d.verbatimMode++
}

// Handle HTML <pre> tag (end)
func (d *Dialectizer) endPre() {
	d.verbatimMode--
}

// Handle the text block processing
func (d *Dialectizer) handleData(text string) {
	if d.verbatimMode > 0 {
		d.output.WriteString(text)
	} else {
		d.output.WriteString(d.process(text))
	}
}

// ChefDialectizer struct inherits from Dialectizer
type ChefDialectizer struct {
	Dialectizer
}

// FuddDialectizer struct inherits from Dialectizer
type FuddDialectizer struct {
	Dialectizer
}

// OldeDialectizer struct inherits from Dialectizer
type OldeDialectizer struct {
	Dialectizer
}

// NewChefDialectizer creates a new instance of ChefDialectizer
func NewChefDialectizer() *ChefDialectizer {
	return &ChefDialectizer{
		Dialectizer{
			subs: [][2]string{
				{`a([nu])`, `u$1`}, {`A([nu])`, `U$1`},
				{`a\B`, `e`}, {`A\B`, `E`}, {`en\b`, `ee`},
				{`\Bew`, `oo`}, {`\Be\b`, `e-a`}, {`\be`, `i`},
				{`\bE`, `I`}, {`\Bf`, `ff`}, {`\Bir`, `ur`},
				{`(\w*?)i(\w*?)$`, `$1ee$2`}, {`\bow`, `oo`},
				{`\bo`, `oo`}, {`\bO`, `Oo`}, {`the`, `zee`},
				{`The`, `Zee`}, {`th\b`, `t`}, {`\Btion`, `shun`},
				{`\Bu`, `oo`}, {`\BU`, `Oo`}, {`v`, `f`},
				{`V`, `F`}, {`w`, `w`}, {`W`, `W`},
				{`([a-z])[.]`, `$1.  Bork Bork Bork!`},
			},
		},
	}
}

// Translate function for processing HTML content with a specific dialect
func translate(url string, dialectName string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer response.Body.Close()

	htmlSource, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}

	var parser Dialectizer
	switch strings.ToLower(dialectName) {
	case "chef":
		parser = NewChefDialectizer().Dialectizer
		// Add cases for FuddDialectizer or OldeDialectizer similarly
	}

	parser.handleData(string(htmlSource))
	return parser.output.String()
}

// Main test function to test all dialects
func test(url string) {
	dialects := []string{"chef"}
	for _, dialect := range dialects {
		output := translate(url, dialect)
		fmt.Printf("Translation in %s:\n%s\n", dialect, output)
	}
}

func main() {
	test("http://diveintopython.org/odbchelper_list.html")
}

