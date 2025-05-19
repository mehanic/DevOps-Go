package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	name, birthyear, mail, country := CallMultireturn()

	fmt.Println("\n-- User Info Summary --")
	fmt.Printf("Name      : %s\n", name)
	fmt.Printf("Birthyear : %d\n", birthyear)
	fmt.Printf("Email     : %s\n", mail)
	fmt.Printf("Country   : %s\n", country)
}

func CallMultireturn() (string, int, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("What's your mail address? ")
	mail, _ := reader.ReadString('\n')
	mail = strings.TrimSpace(mail)

	fmt.Print("What's your birthyear? ")
	birthStr, _ := reader.ReadString('\n')
	birthStr = strings.TrimSpace(birthStr)
	birthyear, _ := strconv.Atoi(birthStr)

	fmt.Print("Where are you from? ")
	country, _ := reader.ReadString('\n')
	country = strings.TrimSpace(country)

	return CreateMultireturn(name, birthyear, mail, country)
}

func CreateMultireturn(name string, birthyear int, mail string, country string) (string, int, string, string) {
	return name, birthyear, mail, country
}
