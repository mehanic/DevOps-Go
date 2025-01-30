package main

import (
	"fmt"
	"strconv"
)


func main() {
	ShowConv()
	if err := Strconv(); err != nil {
		panic(err)
	}
	Interfaces()
	CheckType("info")
	Strconv()
}

// --
// ShowConv demonstrates some type conversion
func ShowConv() {
	// int
	var a = 24

	// float 64
	var b = 2.0

	// convert the int to a float64 for this calculation
	c := float64(a) * b
	fmt.Println(c)

	// fmt.Sprintf is a good way to convert to strings
	precision := fmt.Sprintf("%.2f", b)

	// print the value and the type
	fmt.Printf("%s - %T\n", precision, precision)
}

//---

// CheckType will print based on the
// interface type
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("It's an int!")
	default:
		fmt.Println("not sure what it is...")
	}
}

// Interfaces demonstrates casting
// from anonymous interfaces to types
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// manually check an interface
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	// this one should fail
	if _, ok := i.(int); !ok {
		fmt.Println("uh oh! glad we handled this")
	}
}

//---

// Strconv demonstrates some strconv
// functions
func Strconv() error {
	//strconv is a good way to convert to and from strings
	s := "1234"
	// we can specify the base (10) and precision
	// 64 bit
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// lets try hex
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// we can do other useful things like:
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}
