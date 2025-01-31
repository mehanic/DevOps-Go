package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int `json:"user_id, string`
	Username string
	Address  string `json:", omitempty"`
	Company  string `json:"-"`   // Removed `json:"-"` If you want Company to be part of the JSON output,
}

func main() {
	u := &User{
		ID:       30,
		Username: "hoasker",
		Address:  "",
		Company:  "RealZealot",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("Jsong string: %s\n", string(result))
}

// Jsong string: {"ID":30,"Username":"hoasker","Address":""}
