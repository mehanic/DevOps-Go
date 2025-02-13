package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is a tutorial for struct")
	// no inheritance in golang; No super of parent
	reader := bufio.NewReader(os.Stdin)
	mp := make(map[string]int)

	fmt.Println("Enter marks for physics")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSpace(str1)
	pmarks, _ := strconv.ParseInt(str1, 10, 64)
	mp["physics"] = int(pmarks)

	fmt.Println("Enter marks for chemistry")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	cmarks, _ := strconv.ParseInt(str2, 10, 64)
	mp["chemistry"] = int(cmarks)

	fmt.Println("Enter marks for maths")
	str3, _ := reader.ReadString('\n')
	str3 = strings.TrimSpace(str3)
	mmarks, _ := strconv.ParseInt(str3, 10, 64)
	mp["maths"] = int(mmarks)

	st1 := Student{5966, "Rohan", mp}
	fmt.Println("Here are some details about ", st1.Name)
	fmt.Println(st1)

	fmt.Println("Now printing using printf")
	fmt.Printf("%+v", st1)
}

type Student struct {
	RollNo int
	Name   string
	Result map[string]int
}


// This is a tutorial for struct
// Enter marks for physics
// 5
// Enter marks for chemistry
// 4
// Enter marks for maths
// 7
// Here are some details about  Rohan
// {5966 Rohan map[chemistry:4 maths:7 physics:5]}
// Now printing using printf
// {RollNo:5966 Name:Rohan Result:map[chemistry:4 maths:7 physics:5]}
