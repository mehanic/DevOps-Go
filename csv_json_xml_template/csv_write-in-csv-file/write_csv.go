package main

import (
	"encoding/csv"
	"os"
)

func write_to_csv_list() {
	file, err := os.Create("products.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Category", "Name", "Quantity", "Price"})
	writer.Write([]string{"41", "Furnishings", "Office Chair", "10", "85"})
	writer.Write([]string{"20", "Office Supplies", "Pens", "30", "10"})
	writer.Write([]string{"13", "Housekeeping", "Bed Sheet (Double)", "15", "20"})
}

func write_to_csv_dictionary() {
	file, err := os.OpenFile("products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "category", "name", "quantity", "price"}
	writer.Write(headers)

	item := map[string]string{
		"id":       "65",
		"category": "maintenance",
		"name":     "ladder",
		"quantity": "33",
		"price":    "50",
	}
	writer.Write([]string{item["id"], item["category"], item["name"], item["quantity"], item["price"]})
}

func main() {
	write_to_csv_dictionary()
}
