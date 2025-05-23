package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	fmt.Println("", currentMonth, "", currentYear)
	caL()
}

func caL() {
	langs := [][]string{
		{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
		{"  ", "  ", "  ", " 1", " 2", " 3", " 4"},
		{" 5", " 6", " 7", " 8", " 9", "10", "11"},
		{"12", "13", "14", "15", "16", "17", "18"},
		{"19", "20", "21", "22", "23", "24", "25"},
		{"26", "27", "28", "29", "30", "  ", "  "}}
	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang, " ")
		}
		fmt.Println()
	}
}
