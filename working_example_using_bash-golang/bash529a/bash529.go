package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Println("")
		return
	}

	filename := os.Args[1]

	// Get the file extension
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".tgz", ".tar.gz":
		// Decompress .tgz or .tar.gz
		runCommand("tar", "zxvf", filename)
	case ".tar.z":
		// Decompress .tar.Z
		runCommand("gunzip", filename)
		runCommand("tar", "xvf", strings.TrimSuffix(filename, ".Z"))
	case ".tar.bz2":
		// Decompress .tar.bz2
		runCommand("tar", "xvjf", filename)
	case ".tar":
		// Decompress .tar
		runCommand("tar", "xvf", filename)
	case ".gz":
		// Decompress .gz
		runCommand("gunzip", filename)
	case ".z":
		// Decompress .Z
		runCommand("gunzip", filename)
	case ".bz2":
		// Decompress .bz2
		runCommand("bunzip2", filename)
	case ".zip":
		// Decompress .zip
		runCommand("unzip", filename)
	default:
		fmt.Printf("ファイルの拡張子が対応していません: %s\n", filename)
	}
}

// runCommand executes a command with the provided arguments
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("コマンド %s の実行中にエラーが発生しました: %v\n", name, err)
	}
}

