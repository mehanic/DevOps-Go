package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s1 := "12"
	//s2 := "s"  also for fail
	s2 := "45"
	n1, e1 := strconv.Atoi(s1)
	if e1 != nil {
		log.Fatal(e1)
	}
	n2, e2 := strconv.Atoi(s2)
	if e2 != nil {
		log.Fatal(e2)
	}
	c := n1 + n2
	s3 := strconv.Itoa(c)
	fmt.Println(s1 + s2 + s3)
	fmt.Println("Program completed")
}

// 124557
// Program completed
