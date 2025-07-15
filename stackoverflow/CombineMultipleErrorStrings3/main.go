package main

import (
    "errors"
    "fmt"
)

// Define custom errors
var (
    ErrIncorrectUsername = errors.New("incorrect username")
    ErrIncorrectPassword = errors.New("incorrect password")
)

// validate checks username and password, and returns an appropriate error
func validate(username, password string) error {
    if username != "admin" {
        return ErrIncorrectUsername
    }
    if password != "1234" {
        return ErrIncorrectPassword
    }
    return nil
}

func main() {
    err := validate("ruster", "4321")

    if err == nil {
        fmt.Println("Login successful!")
        return
    }

    // Detect specific errors
    if errors.Is(err, ErrIncorrectUsername) {
        fmt.Println("Username is incorrect")
    }

    if errors.Is(err, ErrIncorrectPassword) {
        fmt.Println("Password is incorrect")
    }
}

//-----------------------------------------------
type ErrorCollector []error

func (c *ErrorCollector) Collect(e error) { *c = append(*c, e) }

func (c *ErrorCollector) Error() (err string) {
    err = "Collected errors:\n"
    for i, e := range *c {
        err += fmt.Sprintf("\tError %d: %s\n", i, e.Error())
    }

    return err
}

//-----------------------------------------

func main1() {
    collector := new(ErrorCollector)

    for i := 0; i < 10; i++ {
         collector.Collect(errors.New(fmt.Sprintf("%d Error", i)))
    }

    fmt.Println(collector)
}


//----------------------


func main2() {
    var errors []error
    errors = append(errors, fmt.Errorf("error 1"))
    errors = append(errors, fmt.Errorf("error 2"))
    errors = append(errors, fmt.Errorf("and yet another error"))

    // Print all errors as a string slice
    for _, err := range errors {
        fmt.Println(err)
    }

    // Or combine them into a single string
    s := fmt.Sprintf("%v", errors)
    fmt.Println("Combined errors:", s)
}

//--------------------------------------------------
var someErr = errors.New("my error")

func f() error {
    return fmt.Errorf("can't process request: %w", someErr)
}

func main3() {
    err := f()

    if errors.Is(err, someErr) {
        fmt.Println("Detected: my error")
    } else {
        fmt.Println("Some other error")
    }
}

///

func NewJoinedErrors(err1 error, err2 error) JoinedErrors {
    return JoinedErrors{err1: err1, err2: err2}
}

type JoinedErrors struct {
    err1 error
    err2 error
}

func (e JoinedErrors) Error() string {
    return fmt.Sprintf("%s: %s", e.err1, e.err2)
}

func (e JoinedErrors) Unwrap() error {
    return e.err2
}

func (e JoinedErrors) Is(target error) bool {
    return errors.Is(e.err1, target)
}

func main4(){

someErr1:=errors.New("my error 1")
someErr2:=errors.New("my error 2")
err:=NewJoinedErrors(someErr1, someErr2)
// this will be true because 
// (e JoinedErrors) Is(target error) 
// will return true 
if errors.Is(err, someErr1){...}
// this will be true because 
// (e JoinedErrors) Unwrap() error 
// will return err2 
if errors.Is(err, someErr2){...}
}

//


func condenseErrors(errs []error) error {
    switch len(errs) {
    case 0:
        return nil
    case 1:
        return errs[0]
    }
    err := errs[0]
    for _, e := range errs[1:] {
        err = errors.Wrap(err, e.Error())
    }
    return err
}


//


func JoinErrs(errs ...error) error {
    var joinErrsR func(string, int, ...error) error
    joinErrsR = func(soFar string, count int, errs ...error) error {
        if len(errs) == 0 {
            if count == 0 {
                return nil
            }
            return fmt.Errorf(soFar)
        }
        current := errs[0]
        next := errs[1:]
        if current == nil {
            return joinErrsR(soFar, count, next...)
        }
        count++
        if count == 1 {
            return joinErrsR(fmt.Sprintf("%s", current), count, next...)
        } else if count == 2 {
            return joinErrsR(fmt.Sprintf("1: %s\n2: %s", soFar, current), count, next...)
        }
        return joinErrsR(fmt.Sprintf("%s\n%d: %s", soFar, count, current), count, next...)
    }
    return joinErrsR("", 0, errs...)
}


///

