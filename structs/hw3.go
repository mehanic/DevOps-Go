package main

import "fmt"

type contact struct{
    email string
    phone string
}

type Person struct{
    name string
    age string
    contact
}
func main() {
    var Tom=Person{
        name:"Tom",
        age:"12",
        contact: contact{
            email:"am@mail",
            phone: "3232131",
        },
    }
    fmt.Println(Tom.email)
}