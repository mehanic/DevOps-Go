package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "carl"


	fmt.Println(len(name))
}

func main1() {
	
	name := "İnanç"
	fmt.Printf("%q is %d bytes\n", name, len(name))


	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))
}

func main2() {
	msg := os.Args[1]


	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)


}

func main3() {
	

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)
}

func main4() {
	

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
}

func main5() {
	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
}
