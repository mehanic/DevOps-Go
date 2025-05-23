package main

import (
	"log"
	"os"
)

func main() {
	// create and open a temporary file
	f, err := os.CreateTemp("", "tmpfile-") // in Go version older than 1.17 you can use ioutil.TempFile
	if err != nil {
		log.Fatal(err)
	}

	// close and remove the temporary file at the end of the program
	defer f.Close()
	defer os.Remove(f.Name())

	// write data to the temporary file
	data := []byte("abc abc abc")
	if _, err := f.Write(data); err != nil {
		log.Fatal(err)
	}
}

