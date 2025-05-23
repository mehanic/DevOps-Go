package main

import (
	"fmt"
	"strings"
)

func main() {
	paths := []string{"/usr/bin", "/usr/bin/httpd", "/home/users/xyz/weblogic/config.xml"}
	for _, currentPath := range paths {
		fmt.Println("now working on: ", currentPath)
		if strings.Contains(currentPath, "httpd") {
			fmt.Println(currentPath)
			break
		}
	}
	fmt.Println("outside of for loop")
}
