package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD = "2006-01-02"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now.Format(YYYYMMDD))
}

//2025-01-17