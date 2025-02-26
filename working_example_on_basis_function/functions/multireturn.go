package functions

import "fmt"

func CallMultireturn() (string, int, string, string) {

	var name, mail, country string
	var birthyear int

	fmt.Println("What's your name?")

	fmt.Scanln(&name)

	fmt.Println("What's your mail address?")

	fmt.Scanln(&mail)

	fmt.Println("What's your birthyear?")

	fmt.Scanln(&birthyear)

	fmt.Println("Where are you from?")

	fmt.Scanln(&country)

	name, birthyear, mail, country = CreateMultireturn(name, birthyear, mail, country)

	return name, birthyear, mail, country

}

func CreateMultireturn(name string, birthyear int, mail string, country string) (string, int, string, string) {
	return name, birthyear, mail, country
}
