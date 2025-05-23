package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	// "io"
	/*       "math"
	         "sort"
	         "os"
	         "io/ioutil"
	         "strconv"*/)

func main() {
	pleString := "Hello world"

	fmt.Println(strings.Contains(sampleString, "lo"))
	fmt.Println(strings.Index(sampleString, "l"))
	fmt.Println(strings.Count(sampleString, "l"))
	fmt.Println(strings.Replace(sampleString, "l", "x", 3))

	String := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	tOfLetters := []string{"c", "a", "b"}
	t.Strings(listOfLetters)

	fmt.Println("Letters: ", listOfLetters)

	listOfNums := strings.Join([]string{"3", "2", "1"}, ",")

	fmt.Println(listOfNums)

	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)

	file.WriteString(" This is some random text")
	e.Close()

	stream, err := io.ReadFile("sample.txt")
	err != nil {
	g.Fatal(err)


	readString := string(stream)
	.Println(readString)


