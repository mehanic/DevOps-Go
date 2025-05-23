package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ajstarks/svgo"
)

type PopulationData struct {
	CountryName string  `json:"Country Name"`
	Value       string  `json:"Value"`
	Year        string  `json:"Year"`
}

func main() {
	// Load JSON data
	filename := "population_data.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse JSON data
	var popData []PopulationData
	if err := json.Unmarshal(data, &popData); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Build population dictionary
	ccPopulations := make(map[string]int)
	for _, popDict := range popData {
		if popDict.Year == "2010" {
			countryName := popDict.CountryName
			population, _ := strconv.Atoi(popDict.Value) // convert population to integer
			code := getCountryCode(countryName)
			if code != "" {
				ccPopulations[code] = population
			}
		}
	}

	// Group countries by population
	ccPops1 := make(map[string]int)
	ccPops2 := make(map[string]int)
	ccPops3 := make(map[string]int)
	for code, pop := range ccPopulations {
		if pop < 10000000 {
			ccPops1[code] = pop
		} else if pop < 1000000000 {
			ccPops2[code] = pop
		} else {
			ccPops3[code] = pop
		}
	}

	// Print the count of countries in each population level
	fmt.Printf("0-10m: %d\n", len(ccPops1))
	fmt.Printf("10m-1bn: %d\n", len(ccPops2))
	fmt.Printf(">1bn: %d\n", len(ccPops3))

	// Create SVG
	file, err := os.Create("World_population.svg")
	if err != nil {
		log.Fatalf("Error creating SVG file: %v", err)
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(800, 600)
	canvas.Rect(0, 0, 800, 600, "fill:#ffffff") // Background

	// Draw some dummy rectangles as placeholders for countries
	// Note: You will need actual coordinates and country shapes for a real map
	drawRegions(canvas, ccPops1, "red")
	drawRegions(canvas, ccPops2, "blue")
	drawRegions(canvas, ccPops3, "green")

	canvas.End()
	fmt.Println("SVG file created successfully")
}
func drawRegions(canvas *svg.SVG, populations map[string]int, color string) {
	for i := 0; i < len(populations); i++ {
		canvas.Rect(50+i*20, 50, 15, 10, fmt.Sprintf("fill:%s", color))
	}
}

// getCountryCode is a placeholder function
func getCountryCode(countryName string) string {
	// Placeholder for country code retrieval
	// Implement this function based on your needs
	return ""
}

