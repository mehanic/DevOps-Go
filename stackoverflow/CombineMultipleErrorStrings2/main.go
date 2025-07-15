package main

import (
	"errors"
	"fmt"
)
var (
    ErrIncorrectUsername = errors.New("incorrect username")
    ErrIncorrectPassword = errors.New("incorrect password")
)

func main() {
    err := validate("ruster", "4321")
    // You can print multi-line errors
    // Each will be separated by a newline character (\n).
    if err != nil {
        fmt.Println(err)
        // incorrect username
        // incorrect password
    }
}

func validate(username, password string) error {
    var errs []error

    // errors.Join the errors into a single error
    if username != "gopher" {
        errs = append(errs, ErrIncorrectUsername)
    }
    if password != "1234" {
        errs = append(errs, ErrIncorrectPassword)
    }

    // Join returns a single `error`.
    // Underlying, the error contains all the errors we add.
    return errors.Join(errs...)
}