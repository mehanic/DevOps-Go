package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("starting dir:", wd)

	if err := os.Chdir("/"); err != nil {
		fmt.Println(err)
		return
	}

	if wd, err = os.Getwd(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("final dir:", wd)
}


//  └> $ go run wd.go 
// starting dir: /home/free/GOLANG/HandsOnSystemProgrammingWithGo/Chapter04/filepath/wd
// final dir: /