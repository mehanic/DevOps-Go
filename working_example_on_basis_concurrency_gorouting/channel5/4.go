package main

import "fmt"

type Student struct {
	Name string
}

var (
	ch = make(chan *Student)
)

func printMessage(message string) {
	ch <- &Student{message}
}

func main() {
	var st1, st2, st3 *Student
	go printMessage("one")
	st1 = <-ch
	go printMessage("two")
	st2 = <-ch
	go printMessage("three")
	st3 = <-ch
	fmt.Println("student1", st1)
	fmt.Println("student2", st2)
	fmt.Println("student3", st3)
}

// student1 &{one}
// student2 &{two}
// student3 &{three}
