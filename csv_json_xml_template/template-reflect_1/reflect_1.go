package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int
	RealName string `unpack:"-"`
	Login    string
	Flags    int
}

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem()

	fmt.Printf("%T have %d fields:\n", u, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("\tname=%v, type=%v, value=%v, tag=`%v`\n", typeField.Name,
			typeField.Type.Kind(),
			valueField,
			typeField.Tag,
		)
	}
	return nil
}

func main() {
	u := &User{
		ID:       42,
		RealName: "Hoasker",
		Flags:    32,
	}
	err := PrintReflect(u)
	if err != nil {
		panic(err)
	}
}

// *main.User have 4 fields:
// 	name=ID, type=int, value=42, tag=``
// 	name=RealName, type=string, value=Hoasker, tag=`unpack:"-"`
// 	name=Login, type=string, value=, tag=``
// 	name=Flags, type=int, value=32, tag=``
