package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Tutorial for method")
	marks := float32(0.0)
	s := Student{
		RollNo: 123,
		Name:   "XYZ",
		marks:  marks,
	}
	s.GetStatus()
}

type Student struct {
	RollNo int
	Name   string
	marks  float32
}

func (s Student) GetStatus() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter RollNo")
	temp1, _ := reader.ReadString('\n')
	rollNo, _ := strconv.ParseInt(strings.TrimSpace(temp1), 10, 64)
	s.RollNo = int(rollNo)

	fmt.Println("Enter nameo")
	name, _ := reader.ReadString('\n')
	s.Name = name

	fmt.Println("Enter Marks")
	temp2, _ := reader.ReadString('\n')
	marks, _ := strconv.ParseFloat(strings.TrimSpace(temp2), 64)
	s.marks = float32(marks)

	fmt.Printf("New Details: \n %+v", s)

}

// Tutorial for methodEnter RollNo
// 1
// Enter nameo
// start
// Enter Marks
// five
// New Details: 
//  {RollNo:1 Name:start
//  marks:0}
