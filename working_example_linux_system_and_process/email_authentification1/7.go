package main

import "fmt"

func checkEmail(email string) bool {
	isCorrect := false
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			isCorrect = true
			break
		}
	}
	return isCorrect
}

func main() {
	var email string
	fmt.Scanf("%s", &email)
	if checkEmail(email) {
		fmt.Println("ok")
	} else {
		fmt.Println("should contain @")
	}
}

// ╰─λ go run 7.go                                                                                        0 (14.517s) < 20:30:21
// canada
// should contain @

// ╰─λ go run 7.go                                                                                         0 (2.942s) < 20:30:35
// kubecon@gmail.com
// ok
