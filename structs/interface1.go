package main

import "fmt"

func main() {
	config := map[string]interface{}{
		"host": "localhost",
		"port": 8080,
		"user": "rock",
		"pass": "@123k9",
		"tls":  true,
	}
	for key, value := range config {
		fmt.Println(key, " : ", value)
	}

}


// user  :  rock
// pass  :  @123k9
// tls  :  true
// host  :  localhost
// port  :  8080