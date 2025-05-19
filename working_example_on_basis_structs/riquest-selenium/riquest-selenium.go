package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	// Start a new Selenium WebDriver session
	const (
		seleniumPath    = "path/to/selenium-server-standalone.jar" // Update this path
		chromeDriverPath = "path/to/chromedriver"                   // Update this path
		port            = 8080
	)

	// Start the Selenium server
	if err := selenium.StartSeleniumServer(seleniumPath, port); err != nil {
		log.Fatalf("Failed to start Selenium server: %v", err)
	}
	defer selenium.StopSeleniumServer()

	// Start a new WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.Set("goog:chromeOptions", map[string]interface{}{
		"args": []string{"--headless"}, // Run Chrome in headless mode
	})

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}
	defer driver.Quit()

	// Navigate to the page
	if err := driver.Get("http://pythonscraping.com/pages/javascript/ajaxDemo.html"); err != nil {
		log.Fatalf("Failed to load page: %v", err)
	}

	// Wait for the page to load
	time.Sleep(3 * time.Second)

	// Find the element by ID and print its text
	content, err := driver.FindElement(selenium.ByID, "content")
	if err != nil {
		log.Fatalf("Failed to find element: %v", err)
	}

	text, err := content.Text()
	if err != nil {
		log.Fatalf("Failed to get text: %v", err)
	}

	fmt.Println(text)
}

