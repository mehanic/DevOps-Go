package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run a command that should always work (mktemp)
	cmd1 := exec.Command("mktemp")
	err := cmd1.Run()
	mktempRc := 0
	if err != nil {
		mktempRc = 1 // Set to 1 if mktemp fails (though it shouldn't)
	}

	// Run a command that should always fail (mkdir /home/example)
	cmd2 := exec.Command("mkdir", "/home/example")
	err = cmd2.Run()
	mkdirRc := 0
	if err != nil {
		mkdirRc = 1 // Set to 1 if mkdir fails
	}

	// Print the return codes
	fmt.Printf("mktemp returned %d, while mkdir returned %d!\n", mktempRc, mkdirRc)
}

