package main

import "fmt"

// struct
func test() {
	fmt.Println("test")
}

func introduce(name string, nickname string) {
	fmt.Println("myname is", name)
	fmt.Printf("my nickname is %v \n", nickname)
}

func main() {

	var goodnightName string = "goodnight"
	var goodnightNickname string = "เท่ียงคืน"

	munbodName := "munbod"
	munbodNickname := "kati"
	introduce(goodnightName, goodnightNickname)
	introduce(munbodName, munbodNickname)
	test()
}


// myname is goodnight
// my nickname is เท่ียงคืน 
// myname is munbod
// my nickname is kati 
// test
