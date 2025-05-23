package main

import (
  "fmt"
)

type greetings func(string) string

func addGreetings(s string) string {
  return fmt.Sprintf("Learning %s", s)
}

func main() {
  var greet greetings
  greet = addGreetings
  fmt.Println(greet("Golang function as type"))
}
