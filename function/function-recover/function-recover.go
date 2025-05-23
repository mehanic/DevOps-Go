package main

import (
	"errors"
	"fmt"
)

func connectToDB() error {
	return errors.New("error while connecting to db")
}

func main() {
	err := connectToDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to db")
}
