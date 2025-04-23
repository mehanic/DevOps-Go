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

	SearchString()
	ModifyString()
	StringReader()
}

func SearchString() {
	s := "this is a test"

	fmt.Println(strings.Contains(s, "this"))

	fmt.Println(strings.ContainsAny(s, "abc"))

	fmt.Println(strings.HasPrefix(s, "this"))

	fmt.Println(strings.HasSuffix(s, "test"))
}

func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array1111111"

	b := Buffer(rawString)

	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text(), " ")
	}

	return nil
}

func ModifyString() {
	s := "simple string"

	fmt.Println(strings.Split(s, " "))

	fmt.Println(strings.Title(s))

	s = " simple string "
	fmt.Println(strings.TrimSpace(s))

	caser := cases.Title(language.English)
	fmt.Println(caser.String(s)) // "Simple String"
}

func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}

// -----------------------------------------------------------
func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	b = bytes.NewBuffer(rawBytes)

	b = bytes.NewBufferString(rawString)

	return b
}

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
