package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time : ", currentTime)
	formattedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Current time in new format: ", formattedTime)
	createdDate := time.Date(2000, time.January, 28, 7, 28, 0, 0, time.UTC)
	fmt.Println("Created Date and time = ", createdDate.Format("01-02-2006 15:04:05 Monday"))
	time.Sleep(time.Duration(time.Minute))
}

// Current time :  2025-01-16 15:00:44.346942738 +0100 CET m=+0.000018423
// Current time in new format:  01-16-2025 15:00:44 Thursday
// Created Date and time =  01-28-2000 07:28:00 Friday
