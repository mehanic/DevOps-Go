package main

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx"
)

func main() {
	files, _ := os.ReadDir(".")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".xlsx" {
			wb, _ := xlsx.OpenFile(file.Name())
			for _, sheet := range wb.Sheets {
				csvFileName, _ := os.Create(file.Name() + sheet.Name + ".csv")
				csvFile := csv.NewWriter(csvFileName)

				for _, row := range sheet.Rows {
					var rowData []string
					for _, cell := range row.Cells {
						cellData := cell.String()
						rowData = append(rowData, cellData)
					}

					csvFile.Write(rowData)
				}

				csvFile.Flush()
				csvFileName.Close()
			}
		}
	}
}