package main

import (
	"os"
	"path/filepath"
)

func main() {
	dir := os.TempDir()
	tempDir, err := os.MkdirTemp(dir, "example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempDir)
	filename := filepath.Join(tempDir, "example.txt")
	if err := os.WriteFile(filename, []byte("Hello, World!"), 0666); err != nil {
		panic(err)
	}
}
