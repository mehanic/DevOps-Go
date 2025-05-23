package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func main() {
	if err := BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := NullEncoding(); err != nil {
		panic(err)
	}
}



// json that has name but not age
const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// Example is a basic struct with age
// and name fields
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

// BaseEncoding shows encoding and
// decoding with normal types
func BaseEncoding() error {
	e := Example{}

	// note that no age = 0 age
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}

//---

type nullInt64 sql.NullInt64

// ExampleNullInt is the same, but
// uses a sql.NullInt64
type ExampleNullInt struct {
	Age  *nullInt64 `json:"age,omitempty"`
	Name string     `json:"name"`
}

func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

// NullEncoding shows an alternative method
// for dealing with nil/omitted values
func NullEncoding() error {
	e := ExampleNullInt{}

	// note that no means an invalid value
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("nullInt64 Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("nullInt64 Marshal, with age = 0:", string(value))

	return nil
}

//---

// ExamplePointer is the same, but
// uses a *Int
type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

// PointerEncoding shows methods for
// dealing with nil/omitted values
func PointerEncoding() error {

	// note that no age = nil age
	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with age = 0:", string(value))

	return nil
}

// Regular Unmarshal, no age: {Age:0 Name:Aaron}
// Regular Marshal, with no age: {"name":"Aaron"}
// Regular Unmarshal, with age = 0: {Age:0 Name:Aaron}
// Regular Marshal, with age = 0: {"name":"Aaron"}

// Pointer Unmarshal, no age: {Age:<nil> Name:Aaron}
// Pointer Marshal, with no age: {"name":"Aaron"}
// Pointer Unmarshal, with age = 0: {Age:0xc000012310 Name:Aaron}
// Pointer Marshal, with age = 0: {"age":0,"name":"Aaron"}

// nullInt64 Unmarshal, no age: {Age:<nil> Name:Aaron}
// nullInt64 Marshal, with no age: {"name":"Aaron"}
// nullInt64 Unmarshal, with age = 0: {Age:0xc000012420 Name:Aaron}
// nullInt64 Marshal, with age = 0: {"age":0,"name":"Aaron"}
