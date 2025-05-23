package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(wg *sync.WaitGroup, url string) {
	response, error := http.Get(url)

	if error != nil {
		fmt.Printf("%s - site is down", error)
	} else {
		defer response.Body.Close()
		fmt.Printf("%s - status code is - %d \n", url, response.StatusCode)

		if response.StatusCode == 200 {
			bodyBytes, error := io.ReadAll(response.Body)

			_ = error

			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response Body to %s\n", file)

			error = ioutil.WriteFile(file, bodyBytes, 0664)
			if error != nil {
				log.Fatal(error)
			}
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	urls := []string{"http://www.google.com", "http://www.medium.com", "http://www.facebook.com"}

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(&wg, url)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of Goroutines:", runtime.NumGoroutine())

	wg.Wait()
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 13:54:08
// ####################
// ####################
// ####################
// No of Goroutines: 4
// http://www.google.com - status code is - 200
// Writing response Body to www.google.com.txt
// http://www.facebook.com - status code is - 200
// Writing response Body to www.facebook.com.txt
// http://www.medium.com - status code is - 200
// Writing response Body to www.medium.com.txt
