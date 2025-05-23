package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	phone    string
}

var jsonStr = `{"id": "42", "username": "hoasker", "phone": "123"}`

func main() {
	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct:\n\t%#v\n\n", u)

	u.phone = "987654321"
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Json string:\n\t%s\n", string(result))
}

// struct:
// 	&main.User{ID:0, Username:"hoasker", phone:""}

// Json string:
// 	{"ID":0,"Username":"hoasker"}
