package main

import "fmt"

func main() {
	names := [...]string{"Einstein", "Shepard", "Tesla"}
	books := [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

// ["Einstein" "Shepard" "Tesla"]
// ["Kafka's Revenge" "Stay Golden" "" "" ""]
