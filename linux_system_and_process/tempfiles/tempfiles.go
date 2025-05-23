package main

import (
	"fmt"
	"os"
)

func main() {
	if err := WorkWithTemp(); err != nil {
		panic(err)
	}
}

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	// If you need a temporary place to store files with the
	// same name ie. template1-10.html a temp directory is a good
	// way to approach it, the first argument being blank means it
	// will use create the directory in the location returned by
	// os.TempDir()
	t, err := os.MkdirTemp("", "tmp")
	if err != nil {
		return err
	}

	// This will delete everything inside the temp directory when this
	// function exits if you want to do this later, be sure to return
	// the directory name to the calling function
	defer os.RemoveAll(t)

	// the directory must exist to create the tempfile
	// t is now a string representing the temporary directory path
	tf, err := os.CreateTemp(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// normally we'd delete the temporary file here, but because
	// we're placing it in a temp directory, it gets cleaned up
	// by the earlier defer

	return nil
}
