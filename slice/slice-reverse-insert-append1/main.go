package main

import "fmt"

func reverseSlice1(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	names := []string{"Aatrox", "Ahri", "Akali", "Alistar"}
	fmt.Println("Before reverse:", names)

	reverseSlice1(names)

	fmt.Println("After reverse:", names)
}
