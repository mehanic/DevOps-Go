package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	realEx, err := filepath.EvalSymlinks(ex)
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(realEx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
}

