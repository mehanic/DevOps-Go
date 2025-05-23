package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func displayCSVReader() {
	f, err := os.Open("monsters.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record[1])
	}
}

func displayCSVReaderDict() {
	f, err := os.Open("monsters.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}

	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record[headerMap["monsterName"]] + " is priced at " + record[headerMap["price"]])
	}
}
func displayCSVPandas() {
	f, err := os.Open("monsters.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	fmt.Println(df)
}

func main() {
	displayCSVPandas()
}


// [12x6] DataFrame

//     imageFile monsterName caption                       ...
//  0: monster01 Mingle      Double Trouble                ...
//  1: monster02 Yodel       Yodelay Hee Hoo!              ...
//  2: monster03 Squido      An Eye for Design             ...
//  3: monster04 Spook       Safe and Sound                ...
//  4: monster05 Melville    Networking Guru               ...
//  5: monster06 Filo        Baker by Day, Techie by Night ...
//  6: monster07 Blade       Monster APPetite              ...
//  7: monster08 Timber      Database Expert               ...
//  8: monster09 Skedaddle   Game of Life                  ...
//  9: monster10 Smiley      Don’t Worry, Be Happy!        ...
//     ...       ...         ...                           ...
//     <string>  <string>    <string>                      ...

// Not Showing: description <string>, price <float>, scariness <int>

