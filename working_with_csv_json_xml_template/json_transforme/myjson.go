package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Tutorial for JSON")
	s := []course{
		{"maths", 99, "abcd1234", []string{"geometry", "quadratic"}},
		{"english", 89, "ab1234", []string{"poem", "grammar"}},
		{"physics", 199, "abcd1234", nil},
	}
	content, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(content))
	decodeJSON()
}

func decodeJSON() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "maths",
                "price": 99,
                "tags": [
                        "geometry",
                        "quadratic"
                ]
        }
	`)
	check := json.Valid(jsonDataFromWeb)
	if check {
		fmt.Println("JSON is valid")
		var c1 course
		json.Unmarshal(jsonDataFromWeb, &c1)
		fmt.Printf("\n#v", c1)
	}
}


// Tutorial for JSON
// [
// 	{
// 		"coursename": "maths",
// 		"price": 99,
// 		"tags": [
// 			"geometry",
// 			"quadratic"
// 		]
// 	},
// 	{
// 		"coursename": "english",
// 		"price": 89,
// 		"tags": [
// 			"poem",
// 			"grammar"
// 		]
// 	},
// 	{
// 		"coursename": "physics",
// 		"price": 199
// 	}
// ]
// JSON is valid

// #v%!(EXTRA main.course={maths 99  [geometry quadratic]})⏎     