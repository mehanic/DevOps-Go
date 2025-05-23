package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func greetUsers(names []string) {
	titleCaser := cases.Title(language.English) // Создаём преобразователь регистра
	for _, name := range names {
		msg := "Hello, " + titleCaser.String(name) + "!"
		fmt.Println(msg)
	}
}

func main() {
	usernames := []string{"hannah", "ty", "margot"}
	greetUsers(usernames)
}
