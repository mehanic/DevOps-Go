package strings

import (
	"fmt"
	s "strings"
)

func CreateString() {

	message := "We learn Golang!"

	// it counts "e" in the message
	fmt.Printf("We found e in the message %d times\n", s.Count(message, "e"))

	// it looks if the message contains r
	fmt.Println("Does message contains r? : ", s.Contains(message, "r"))

	// it looks index of the first instance of a in the message
	fmt.Printf("The first Index of the a is %v\n", s.Index(message, "a"))

	// it converts message to the lower case
	fmt.Println(s.ToLower(message))

	// it converts message to the upper case
	fmt.Println(s.ToUpper(message))

}

func StringAdvanced() {

	message := "Learning a new language can be harder sometimes but it makes our minds widen every time!"

	// it looks if message begins with "Learning" suffix
	fmt.Println(s.HasPrefix(message, "Learning"))

	// it looks if message ends with "every" suffix
	fmt.Println(s.HasSuffix(message, "every"))

	languages := []string{"Golang", "Python", "Java", "C", "Javascript", "Typescript"}

	// it joins all the languages in the array with ","

	known_languages := s.Join(languages, ",")
	fmt.Println(known_languages)

	// it replaces "," with "-" 2 times
	fmt.Println(s.Replace(known_languages, ",", "-", 2))

	// it replaces "," with "-" for all of them
	fmt.Println(s.ReplaceAll(known_languages, ",", "-"))

	// it splits variables by ","
	fmt.Println(s.Split(known_languages, ","))

	// it returns string repeatly.
	fmt.Println(s.Repeat(known_languages, 3))

}
