package strings

import (
	"fmt"
	s "strings"
)

func StringLen(data string) int {
	return len(data)
}

func StringToUpper(data string) string {
	return s.ToUpper(data)
}

func StringSplit(data string, delim string) []string {
	return s.Split(data, delim)
}

func StringReplace() {
	var name = "Amsterdam"
	fmt.Println("New string after replacing:", s.Replace(name, "R", "M", 1))
	fmt.Println("New string after replacing:", s.Replace(name, "a", "k", 2))
}

func StringJoin(elements []string) string {
	return s.Join(elements, " ")
}
