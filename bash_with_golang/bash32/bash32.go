package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func waitForLoad(driver selenium.WebDriver) {
	// Find the HTML element
	elem, err := driver.FindElement(selenium.ByTagName, "html")
	if err != nil {
		log.Println("Error finding element:", err)
		return
	}
	count := 0

	for {
		count++
		if count > 20 {
			fmt.Println("Timing out after 10 seconds and returning")
			return
		}

		// Pause for 0.5 seconds
		time.Sleep(500 * time.Millisecond)

		// Check if the page has been reloaded by comparing the html element again
		_, err = driver.FindElement(selenium.ByTagName, "html")
		if err != nil {
			// If an error occurs, it means the page has reloaded
			if _, stale := err.(*selenium.StaleElementReferenceError); stale {
				return
			}
			log.Println("Error during waiting for page load:", err)
			return
		}
	}
}

func main() {
	// Path to the WebDriver (ChromeDriver, GeckoDriver, etc.)
	const (
		seleniumPath = "path/to/selenium-server-standalone.jar"
		driverPath   = "path/to/chromedriver" // Change to your WebDriver path
		port         = 8080
	)

	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(driverPath, port, opts...)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver instance: %v", err)
	}
	defer driver.Quit()

	// Load the desired webpage
	if err := driver.Get("http://pythonscraping.com/pages/javascript/redirectDemo1.html"); err != nil {
		log.Fatalf("Error loading page: %v", err)
	}

	// Wait for the page to fully load
	waitForLoad(driver)

	// Print the page source
	pageSource, err := driver.PageSource()
	if err != nil {
		log.Fatalf("Error getting page source: %v", err)
	}
	fmt.Println(pageSource)
}

