package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study with GOLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("Monday 01-02-2006 15:04:05"))

	createdDate := time.Date(1996, time.July, 30, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday 01-02-2006 15:04:05"))

}


// Welcome to time study with GOLang
// 2025-01-16 14:58:57.317207992 +0100 CET m=+0.000024199
// Thursday 01-16-2025 14:58:57
// 1996-07-30 23:23:00 +0000 UTC
