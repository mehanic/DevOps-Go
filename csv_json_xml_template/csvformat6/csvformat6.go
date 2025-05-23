package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Open the CSV file
	filename := "death_valley_2014.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read CSV data
	reader := csv.NewReader(file)
	headerRow, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading CSV header: %v", err)
	}
	fmt.Println("Header:", headerRow)

	var dates []time.Time
	var highs []float64
	var lows []float64

	// Parse CSV data
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			log.Println("Skipping row due to error:", err)
			continue
		}

		currentDate, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			log.Println("Skipping row due to date parsing error:", err)
			continue
		}
		high, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Println("Skipping row due to high temperature parsing error:", err)
			continue
		}
		low, err := strconv.ParseFloat(row[2], 64) // Adjusted column index
		if err != nil {
			log.Println("Skipping row due to low temperature parsing error:", err)
			continue
		}

		dates = append(dates, currentDate)
		highs = append(highs, high)
		lows = append(lows, low)
	}

	// Create a new plot
	p := plot.New()

	// Create high and low temperature lines
	highsLine, lowsLine := makePlotData(dates, highs, lows)

	// Convert to line plots
	highLine, err := plotter.NewLine(highsLine)
	if err != nil {
		log.Fatalf("Error creating high temperature line: %v", err)
	}

	lowLine, err := plotter.NewLine(lowsLine)
	if err != nil {
		log.Fatalf("Error creating low temperature line: %v", err)
	}

	// Add lines to plot
	p.Add(highLine, lowLine)

	// Add a shaded region between high and low temperatures
	shaded, err := plotter.NewPolygon(createPolygon(dates, highs, lows))
	if err != nil {
		log.Fatalf("Error creating shaded region: %v", err)
	}
	shaded.Color = plotutil.Color(1) // Fix FillColor issue
	p.Add(shaded)

	// Format plot
	p.Title.Text = "Daily High and Low Temperatures - 2014\nDeath Valley, CA"
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Temperature (F)"
	p.X.Tick.Label.Rotation = 45

	// Save plot to PNG file
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "temperature_plot.png"); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}
	fmt.Println("Plot saved to temperature_plot.png")
}

// Helper function to create plot data
func makePlotData(dates []time.Time, highs, lows []float64) (plotter.XYs, plotter.XYs) {
	var highXYs, lowXYs plotter.XYs
	for i, _ := range dates {
		highXYs = append(highXYs, plotter.XY{X: float64(i), Y: highs[i]})
		lowXYs = append(lowXYs, plotter.XY{X: float64(i), Y: lows[i]})
	}
	return highXYs, lowXYs
}

// Helper function to create polygon data
func createPolygon(dates []time.Time, highs, lows []float64) plotter.XYs {
	var points plotter.XYs
	for i := 0; i < len(dates); i++ {
		points = append(points, plotter.XY{X: float64(i), Y: highs[i]})
	}
	for i := len(dates) - 1; i >= 0; i-- {
		points = append(points, plotter.XY{X: float64(i), Y: lows[i]})
	}
	return points
}
