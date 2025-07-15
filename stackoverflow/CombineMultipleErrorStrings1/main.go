package main

import (
    "fmt"
    "strings"
    "syscall"
)

func main() {

    // create a slice for the errors
    var errstrings []string 

    // first error
    err1 := fmt.Errorf("First error:server error")
    errstrings = append(errstrings, err1.Error())

    // do something 
    err2 := fmt.Errorf("Second error:%s", syscall.ENOPKG.Error())
    errstrings = append(errstrings, err2.Error())

    // do something else
    err3 := fmt.Errorf("Third error:%s", syscall.ENOTCONN.Error())
    errstrings = append(errstrings, err3.Error())

    // combine and print all the error
    fmt.Println(fmt.Errorf(strings.Join(errstrings, "\n")))


}