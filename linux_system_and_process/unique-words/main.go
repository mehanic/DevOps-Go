package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++
		words[in.Text()]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}


// go run main.go < shakespeare.txt                       
// There are 99 words, 70 of them are unique.