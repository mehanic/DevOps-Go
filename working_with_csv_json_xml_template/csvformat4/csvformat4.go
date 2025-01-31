package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Series struct {
	Series string  `json:"series"`
	Data   []Point `json:"data"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("anscombe_raw.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Parse the CSV into the structured data
	seriesData := map[string][]Point{
		"I":   {},
		"II":  {},
		"III": {},
		"IV":  {},
	}

	for _, row := range records[2:] { // Skipping the first two header rows
		// Helper function to parse X, Y values
		parseXY := func(xCol, yCol int) (Point, error) {
			x, err := strconv.ParseFloat(row[xCol], 64)
			if err != nil {
				return Point{}, err
			}
			y, err := strconv.ParseFloat(row[yCol], 64)
			if err != nil {
				return Point{}, err
			}
			return Point{X: x, Y: y}, nil
		}

		// Append the data to each series
		for _, series := range []struct {
			name  string
			xCol  int
			yCol  int
		}{
			{"I", 1, 2},
			{"II", 3, 4},
			{"III", 5, 6},
			{"IV", 7, 8},
		} {
			point, err := parseXY(series.xCol, series.yCol)
			if err != nil {
				fmt.Println("Error parsing values:", err)
				return
			}
			seriesData[series.name] = append(seriesData[series.name], point)
		}
	}

	// Create the JSON structure
	var document []Series
	for _, series := range []string{"I", "II", "III", "IV"} {
		document = append(document, Series{
			Series: series,
			Data:   seriesData[series],
		})
	}

	// Print the JSON
	jsonData, err := json.MarshalIndent(document, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

// [
//   {
//     "series": "I",
//     "data": [
//       {
//         "x": 10,
//         "y": 8.04
//       },
//       {
//         "x": 8,
//         "y": 6.95
//       },
//       {
//         "x": 13,
//         "y": 7.58
//       },
//       {
//         "x": 9,
//         "y": 8.81
//       },
//       {
//         "x": 11,
//         "y": 8.33
//       },
//       {
//         "x": 14,
//         "y": 9.96
//       },
//       {
//         "x": 6,
//         "y": 7.24
//       },
//       {
//         "x": 4,
//         "y": 4.26
//       },
//       {
//         "x": 12,
//         "y": 10.84
//       },
//       {
//         "x": 7,
//         "y": 4.82
//       },
//       {
//         "x": 5,
//         "y": 5.68
//       }
//     ]
//   },
//   {
//     "series": "II",
//     "data": [
//       {
//         "x": 10,
//         "y": 9.14
//       },
//       {
//         "x": 8,
//         "y": 8.14
//       },
//       {
//         "x": 13,
//         "y": 8.74
//       },
//       {
//         "x": 9,
//         "y": 8.77
//       },
//       {
//         "x": 11,
//         "y": 9.26
//       },
//       {
//         "x": 14,
//         "y": 8.1
//       },
//       {
//         "x": 6,
//         "y": 6.13
//       },
//       {
//         "x": 4,
//         "y": 3.1
//       },
//       {
//         "x": 12,
//         "y": 9.13
//       },
//       {
//         "x": 7,
//         "y": 7.26
//       },
//       {
//         "x": 5,
//         "y": 4.74
//       }
//     ]
//   },
//   {
//     "series": "III",
//     "data": [
//       {
//         "x": 10,
//         "y": 7.46
//       },
//       {
//         "x": 8,
//         "y": 6.77
//       },
//       {
//         "x": 13,
//         "y": 12.74
//       },
//       {
//         "x": 9,
//         "y": 7.11
//       },
//       {
//         "x": 11,
//         "y": 7.81
//       },
//       {
//         "x": 14,
//         "y": 8.84
//       },
//       {
//         "x": 6,
//         "y": 6.08
//       },
//       {
//         "x": 4,
//         "y": 5.39
//       },
//       {
//         "x": 12,
//         "y": 8.15
//       },
//       {
//         "x": 7,
//         "y": 6.42
//       },
//       {
//         "x": 5,
//         "y": 5.73
//       }
//     ]
//   },
//   {
//     "series": "IV",
//     "data": [
//       {
//         "x": 8,
//         "y": 6.58
//       },
//       {
//         "x": 8,
//         "y": 5.76
//       },
//       {
//         "x": 8,
//         "y": 7.71
//       },
//       {
//         "x": 8,
//         "y": 8.84
//       },
//       {
//         "x": 8,
//         "y": 8.47
//       },
//       {
//         "x": 8,
//         "y": 7.04
//       },
//       {
//         "x": 8,
//         "y": 5.25
//       },
//       {
//         "x": 19,
//         "y": 12.5
//       },
//       {
//         "x": 8,
//         "y": 5.56
//       },
//       {
//         "x": 8,
//         "y": 7.91
//       },
//       {
//         "x": 8,
//         "y": 6.89
//       }
//     ]
//   }
// ]
