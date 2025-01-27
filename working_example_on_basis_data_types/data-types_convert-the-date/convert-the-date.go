package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	s1 := t.String()
	fmt.Printf("s1: %s\n", s1)

	s2 := t.Format("2006-01-02")
	fmt.Printf("s2: %s\n", s2)
}

