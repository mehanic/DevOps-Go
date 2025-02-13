
package main

import "fmt"

type age struct {
    amount int
}

type person struct {
    name string
    age
}

func (a age) tellAge() string {
    return fmt.Sprintf("my age %v", a.amount)
}

type speaker interface {
    tellAge() string
}

func main() {

    a := age{30}
    p := person{"Max", a}

    fmt.Println(p)

    p2 := person{
        "Maria",
        age{26},
    }

    fmt.Println(p2)

    fmt.Printf("name: %v, age: %v\n", p.name, p.amount)

    fmt.Println("also age:", p.age.amount)

    fmt.Println("p.tellAge():", p.tellAge())

    var s speaker = p2
    fmt.Println("speaker:", s.tellAge())

}

// {Max {30}}
// {Maria {26}}
// name: Max, age: 30
// also age: 30
// p.tellAge(): my age 30
// speaker: my age 26

