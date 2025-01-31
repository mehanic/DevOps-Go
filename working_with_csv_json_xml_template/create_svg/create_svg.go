package main

import (
	"log"
	"os"

	"github.com/ajstarks/svgo"
)

func main() {
	// Create an SVG file
	file, err := os.Create("map.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create SVG canvas
	canvas := svg.New(file)
	canvas.Start(800, 600)

	// Title
	canvas.Text(400, 20, "North, Central, and South America", "text-anchor:middle")

	// Draw regions - This is simplified; you would need to add actual map data
	// Define colors for each region
	colors := map[string]string{
		"North America":  "#FF9999",
		"Central America": "#66B2FF",
		"South America": "#99FF99",
		"China":          "#FFCC99",
		"Japan":          "#FFCCFF",
		"Australia":      "#FFFF99",
		"New Zealand":    "#FF9999",
		"Antarctica":     "#CCCCCC",
		"India":          "#FF6666",
	}

	// Sample regions (note: these are not actual coordinates)
	// In a real application, you would need to define actual coordinates for each country
	regions := map[string][][4]int{
		"North America":  {{100, 100, 200, 200}},
		"Central America": {{200, 200, 300, 300}},
		"South America": {{300, 300, 400, 400}},
		"China":          {{400, 400, 500, 500}},
		"Japan":          {{500, 500, 600, 600}},
		"Australia":      {{600, 600, 700, 700}},
		"New Zealand":    {{700, 700, 800, 800}},
		"Antarctica":     {{0, 0, 100, 100}},
		"India":          {{100, 100, 200, 200}},
	}

	// Draw each region with its color
	for region, coords := range regions {
		for _, coord := range coords {
			canvas.Rect(coord[0], coord[1], coord[2]-coord[0], coord[3]-coord[1], "fill:"+colors[region])
		}
	}

	// End SVG canvas
	canvas.End()

	log.Println("SVG file created successfully")
}
