package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for i := 0; i < 3; i++ {
		fmt.Fprintln(w, "hello")
		log.Println(i)
		time.Sleep(time.Second)
	}
}



// └> $ go run file.go 
// 2024/02/25 00:21:05 0
// 2024/02/25 00:21:06 1
// 2024/02/25 00:21:07 2
