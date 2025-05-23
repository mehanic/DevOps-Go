package main

import (
	"fmt"
	"os"
)

func main() {
	cmd, err := parseArgs()
	handleError(err)
	handleError(execCommand(cmd, os.Stdout, os.Stderr))
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
