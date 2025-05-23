package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)


func main() {
	if err := Operate(); err != nil {
		panic(err)
	}

	if err := CapitalizerExample(); err != nil {
		panic(err)
	}

}

//----

// Operate manipulates files and directories
func Operate() error {
	// this 0755 is similar to what you'd see with chown
	// on a command line this will create a director /tmp/example,
	// you may also use an absolute path instead of a relative one
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}

	// go to the /tmp directory
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f is a generic file object
	// it also implements multiple interfaces
	// and can be used as a reader or writer
	// if the correct bits are set when opening
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// we write a known-length value to the file and validate
	// that it wrote correctly
	value := []byte("hello\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// read the file
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// go to the /tmp directory
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// cleanup, os.RemoveAll can be dangerous if you
	// point at the wrong directory, use user input,
	// and especially if you run as root
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}

//----

// Capitalizer opens a file, reads the contents,
// then writes those contents to a second file
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, io.SeekStart); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}

	s := strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}

// CapitalizerExample creates two files, writes to one
// then calls Capitalizer() on both
func CapitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte(`
    this file contains
    a number of words
    and new lines`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}
