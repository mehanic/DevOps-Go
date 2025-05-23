package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToNumber(str string) int {
  // var x int
  x, err := strconv.Atoi(str)
  if err != nil {
    fmt.Println(err)

  }
  return x
}

func main() {
  fmt.Println(StringToNumber("123"))
  main1()
}



func DNAStrand(dna string) string {
  dnaMap := map[string]string{
    "A": "T",
    "T": "A",
    "C": "G",
    "G": "C",
  }
  var result strings.Builder

  // fmt.Println(dnaMap["A"])
  // fmt.Println(dna)
  parts := strings.Split(dna, "")
  for _, s := range parts {
    if s == "A" {
      result.WriteString(dnaMap["A"])
    } else if s == "T" {
      result.WriteString(dnaMap["T"])

    } else if s == "C" {
      result.WriteString(dnaMap["C"])

    } else if s == "G" {
      result.WriteString(dnaMap["G"])
    }
    // fmt.Println(result.String())

  }
  return result.String()

}

func main1() {

  fmt.Println(DNAStrand("CATA"))

}
