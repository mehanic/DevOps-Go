package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Step 1: Create 1.txt and 2.txt
	err := ioutil.WriteFile("1.txt", []byte("1\n2\n3\n4\n"), 0644)
	if err != nil {
		log.Fatalf("Failed to create 1.txt: %v", err)
	}
	err = ioutil.WriteFile("2.txt", []byte("1\n3\n4\n"), 0644)
	if err != nil {
		log.Fatalf("Failed to create 2.txt: %v", err)
	}

	// Step 2: Use 'diff' to find the difference between the files
	cmdDiff := exec.Command("diff", "-u", "1.txt", "2.txt")
	patchFile, err := os.Create("12.patch")
	if err != nil {
		log.Fatalf("Failed to create 12.patch: %v", err)
	}
	defer patchFile.Close()

	cmdDiff.Stdout = patchFile
	if err := cmdDiff.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() != 1 {
			log.Fatalf("Failed to run diff: %v", err)
		}
	}

	fmt.Println("Diff file generated: 12.patch")

	// Step 3: Apply the patch using the 'patch' command
	cmdPatch := exec.Command("patch", "1.txt", "-i", "12.patch")
	if err := cmdPatch.Run(); err != nil {
		log.Fatalf("Failed to apply patch: %v", err)
	}

	fmt.Println("Patch applied to 1.txt")
}

