package main

import (
	"fmt"
	"log"
	"os/exec"
)

func CaptureOutput() {

	out, err := exec.Command("ls", "-l").Output() //  просто сразу возвращаем вывод команды

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func main() {
	CaptureOutput()
}

//show like command ls -la
