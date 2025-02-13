
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)



type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// implementation of the sort.Interface:

// by default `list` sorts by `Title`.
func (l list) Len() int           { return len(l) }
func (l list) Less(i, j int) bool { return l[i].Title < l[j].Title }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

// byRelease sorts by product release dates.
type byRelease struct{ list }

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}




type money float64

func (m money) String() string {
	return fmt.Sprintf("$%.2f", m)
}


type product struct {
	Title    string    `json:"title"`
	Price    money     `json:"price"`
	Released timestamp `json:"released"`
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}


type timestamp struct {
	time.Time
}

// implementation of the fmt.Stringer interface:

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"
	return ts.Format(layout)
}

// implementation of the json.Marshaler and json.Unmarshaler interfaces:

// timestamp knows how to decode itself from json.
// json.Unmarshal and json.Decode call this method.
func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}

// timestamp knows how to encode itself to json.
// json.Marshal and json.Encode call this method.
func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}





const data = `[
 {
  "title": "moby dick",
  "price": 10,
  "released": 118281600
 },
 {
  "title": "odyssey",
  "price": 15,
  "released": 733622400
 },
 {
  "title": "hobbit",
  "price": 25,
  "released": -62135596800
 }
]`

func main() {
	/* encoding */
	l := list{
		{Title: "moby dick", Price: 10, Released: toTimestamp(118281600)},
		{Title: "odyssey", Price: 15, Released: toTimestamp("733622400")},
		{Title: "hobbit", Price: 25},
	}

	data, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	/* decoding */
	// var l list
	// if err := json.Unmarshal([]byte(data), &l); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(l)
}


// [
//  {
//   "title": "moby dick",
//   "price": 10,
//   "released": 118281600
//  },
//  {
//   "title": "odyssey",
//   "price": 15,
//   "released": 733622400
//  },
//  {
//   "title": "hobbit",
//   "price": 25,
//   "released": -62135596800
//  }
// ]

