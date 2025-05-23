package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	//"io/ioutil"
)

//	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter3/encoding"

func main() {
	if err := Base64Example(); err != nil {
		panic(err)
	}

	if err := Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := GobExample(); err != nil {
		panic(err)
	}
}

//---

// Base64Example demonstrates using
// the base64 package
func Base64Example() error {
	// base64 is useful for cases where
	// you can't support binary formats
	// it operates on bytes/strings

	// using helper functions and URL encoding
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	// decode the first value
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding: ", string(decoded))

	return nil
}

// Base64ExampleEncoder shows similar examples
// with encoders/decoders
func Base64ExampleEncoder() error {
	// using encoder/ decoder
	buffer := bytes.Buffer{}

	// encode into the buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)

	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}

	// be sure to close
	if err := encoder.Close(); err != nil {
		return err
	}

	fmt.Println("Using encoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}

	fmt.Println("Using decoder and StdEncoding: ", string(results))

	return nil
}

// ------
// pos stores the x, y position
// for Object
type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample demonstrates using
// the gob package
func GobExample() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// note that if p was an interface
	// we'd have to call gob.Register first

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	// note this is a binary format so it wont print well
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}

// With EncodeToString and URLEncoding:  ZW5jb2Rpbmcgc29tZSBkYXRhIQ==
// With DecodeToString and URLEncoding:  encoding some data!
// Using encoder and StdEncoding:  ZW5jb2Rpbmcgc29tZSBvdGhlciBkYXRh
// Using decoder and StdEncoding:  encoding some other data
// Gob Encoded valued length:  56
// Gob Decode value:  {10 15 wrench}
