package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	err := WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	SearchString()
	ModifyString()
	StringReader()
}

// SearchString shows a number of methods
// for searching a string
func SearchString() {
	s := "this is a test"

	// returns true because s contains
	// the word this
	fmt.Println(strings.Contains(s, "this"))

	// returns true because s contains the letter a
	// would also match if it contained b or c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// returns true because s starts with this
	fmt.Println(strings.HasPrefix(s, "this"))

	// returns true because s ends with test
	fmt.Println(strings.HasSuffix(s, "test"))
}

// WorkWithBuffer will make use of the buffer created by the
// Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array1111111"

	b := Buffer(rawString)

	// we can quickly convert a buffer back into bytes with
	// b.Bytes() or a string with b.String()
	fmt.Println(b.String())

	// because this is an io Reader, we can make use of generic
	// io reader functions such as
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// we can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner
	// interfaces
	reader := bytes.NewReader([]byte(rawString))

	// we can also plug it into a scanner that allows buffered
	// reading and tokenization
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterate over all of the scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text(), " ")
	}

	return nil
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"

	// prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// prints "Simple String" (capitalizes the first letter of each word)
	fmt.Println(strings.Title(s))

	// prints "simple string"; all trailing and
	// leading white space is removed
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))

	// Demonstrate title case using x/text/cases for more robust handling
	caser := cases.Title(language.English)
	fmt.Println(caser.String(s)) // "Simple String"
}

// StringReader demonstrates how to create
// an io.Reader interface quickly with a string
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// prints s on Stdout
	io.Copy(os.Stdout, r)
}

// -----------------------------------------------------------
// Buffer function creates and returns a new buffer from raw string data
func Buffer(rawString string) *bytes.Buffer {
	// We'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)

	// There are a number of ways to create a buffer from the
	// raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// Alternatively
	b = bytes.NewBuffer(rawBytes)

	// And avoiding the initial byte array altogether
	b = bytes.NewBufferString(rawString)

	return b
}

// ToString is an example of taking an io.Reader and consuming
// it all, then returning a string
func toString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}


// it's easy to encode unicode into a byte array1111111
// it's easy to encode unicode into a byte array1111111
// it's easy to encode unicode into a byte array1111111 true
// true
// true
// true
// [simple string]
// Simple String
// simple string
//  Simple String 
// simple string
