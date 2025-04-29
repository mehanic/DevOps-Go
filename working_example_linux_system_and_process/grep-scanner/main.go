package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
//This code checks if there is exactly one command-line argument 
//provided (the pattern to search for) and assigns it to the variable pattern.
	var pattern string
	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}


// go run main.go night  < shakespeare.txt                       
// come night come romeo come thou day in night
// for thou wilt lie upon the wings of night
// come gentle night come loving black-browed night
// that all the world will be in love with night
// ┌─(~/structure/system_i_process/04-grep)──────

//go run main.go yet  < shakespeare.txt                   
//not yet enjoyed
