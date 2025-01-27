package main

import "fmt"

func Greeting(time int) string {
	if time > 6 && time < 12 {
		return "Good morning"
	}
	if time < 16 {
		return "Good afternoon"
	}
	if time < 20 {
		return "Good evening"
	}
	return "Good night"
}
func hello(name string) string {
	return "my name is " + name
}

func main() {
	fmt.Println("lesson 2")
	goodMorning := Greeting(10)
	goodNight := Greeting(21)
	helloMunbod := hello("munbod")
	helloGoodnight := hello("goodnight")
	fmt.Println(goodMorning, helloMunbod)
	fmt.Println(goodNight, helloGoodnight)
}

// lesson 2
// Good morning my name is munbod
// Good night my name is goodnight
