package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"clement", 29}
	json.NewEncoder(os.Stdout).Encode(p)
}

//{"Name":"clement","Age":29}