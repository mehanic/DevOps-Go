package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var start, end string
	sumi := 0
	fmt.Scan(&start)
	fmt.Scan(&end)
	startInt, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	endInt, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}
	for i := startInt; i <= endInt; i++ {
		sumi = sumi + i
	}
	fmt.Println("cумма числе от начало и до конца", sumi)
	fmt.Print("Вывод четных чисел ")
	for i := startInt; i <= endInt; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	sumi = 0
	fmt.Println(" ")
	for i := startInt; i <= endInt; i++ {
		if i%2 != 0 {
			sumi = sumi + i
		}
	}
	fmt.Println("Вывод суммы нечетных чисел", sumi)
}
