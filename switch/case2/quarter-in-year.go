package main

import "fmt"

func main() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for i, month := range months {
		fmt.Println(i+1, month, quarterInYear(month)) //Có thể dùng fmt.Printf("%d - %s - %s\n", i+1, month, quarterInYear(month))
	}
}

func quarterInYear(month string) string {
	switch month {
	case "Jan", "Feb", "Mar":
		return "1st quarter"
	case "Apr", "May", "Jun":
		return "2nd quarter"
	case "Jul", "Aug", "Sep":
		return "3rd quarter"
	default:
		return "4th quarter"
	}
}
