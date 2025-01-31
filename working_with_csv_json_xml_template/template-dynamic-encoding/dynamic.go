package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `[
	{"id": 17, "username": "hoasker", "phone": 0},
	{"id": "17", "address": "none", "company": "RealZealot"}
]`

func main() {
	data := []byte(jsonStr)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("Unpacked in empty interface:\n%#v\n\n", user1)

	user2 := map[string]interface{}{
		"id":       42,
		"username": "hoasker",
	}
	var user2i interface{} = user2
	result, _ := json.Marshal(user2i)
	fmt.Printf("Json string from map:\n %s\n", string(result))
}

// Unpacked in empty interface:
// []interface {}{map[string]interface {}{"id":17, "phone":0, "username":"hoasker"}, map[string]interface {}{"address":"none", "company":"RealZealot", "id":"17"}}

// Json string from map:
//  {"id":42,"username":"hoasker"}
