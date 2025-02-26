package main

import "fmt"

func main() {
	var line string

	line = join(",", []string{"Sammy", "Jessica", "Drew", "Jamie"})
	fmt.Println(line)

	line = join(",", []string{"Sammy", "Jessica"})
	fmt.Println(line)

	line = join(",", []string{"Sammy"})
	fmt.Println(line)
}

func join(del string, values []string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}

