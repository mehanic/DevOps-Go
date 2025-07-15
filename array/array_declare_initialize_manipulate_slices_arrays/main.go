package main

import (
	"log"
)

func main() {
	// var example - slices
	var a []int
	log.Printf("a type=%T, a length=%d, a capacity=%d, a=%v, nil? %t\n", a, len(a), cap(a), a, a == nil)

	// new example - slices
	b := new([]int)
	log.Printf("b type=%T, b length=%d, b capacity=%d, b=%v, nil? %t\n", b, len(*b), cap(*b), b, b == nil)

	// make examples = slices
	c := make([]int, 5)
	log.Printf("c type=%T, c length=%d, c capacity=%d, c=%v, nil? %t\n", c, len(c), cap(c), c, c == nil)

	d := make([]int, 5, 100)
	log.Printf("d type=%T, d length=%d, d capacity=%d, d=%v, nil? %t\n", d, len(d), cap(d), d, d == nil)
	main1()
	main2()
}

func main1() {
	// var example - array
	var e [3]int
	// log.Printf("e type=%T, e length=%d, e capacity=%d, e=%v, nil? %t\n", e, len(e), cap(e), e, e == nil)
	log.Printf("e type=%T, e length=%d, e capacity=%d, e=%v\n", e, len(e), cap(e), e)

	// new example - array
	f := new([3]int)
	log.Printf("f type=%T, f length=%d, f capacity=%d, f=%v, nil? %t\n", f, len(f), cap(f), f, f == nil)

	// slice literals
	g := []int{1, 3, 6}
	log.Printf("g type=%T, g length=%d, g capacity=%d, g=%v, nil? %t\n", g, len(g), cap(g), g, g == nil)

	h := []int{}
	log.Printf("h type=%T, h length=%d, h capacity=%d, h=%v, nil? %t\n", h, len(h), cap(h), h, h == nil)

	// array literals
	j := [3]int{1, 3, 5}
	log.Printf("j type=%T, j length=%d, j capacity=%d, j=%v\n", j, len(j), cap(j), j)

	k := [...]int{1, 3, 5, 7}
	log.Printf("k type=%T, k length=%d, k capacity=%d, k=%v\n", k, len(k), cap(k), k)
}

type champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}

func main2() {
	i := []champion{
		{
			Name:    "Evelynn",
			Classes: []string{"Assassin"},
			Origins: []string{"Demon"},
			Cost:    3,
		},
		{
			Name:    "Vi",
			Classes: []string{"Brawler"},
			Origins: []string{"Hextech"},
			Cost:    3,
		},
	}
	log.Printf("i type=%T, i length=%d, i capacity=%d, i=%v, nil? %t\n", i, len(i), cap(i), i, i == nil)
}

//:!go run DevOps-Go/working_example_on_basis_array/array_declare_initialize_manipulate_slices_arrays/main.go
//2025/05/11 17:16:36 a type=[]int, a length=0, a capacity=0, a=[], nil? true
//2025/05/11 17:16:36 b type=*[]int, b length=0, b capacity=0, b=&[], nil? false
//2025/05/11 17:16:36 c type=[]int, c length=5, c capacity=5, c=[0 0 0 0 0], nil? false
//2025/05/11 17:16:36 d type=[]int, d length=5, d capacity=100, d=[0 0 0 0 0], nil? false
//2025/05/11 17:16:36 e type=[3]int, e length=3, e capacity=3, e=[0 0 0]
//2025/05/11 17:16:36 f type=*[3]int, f length=3, f capacity=3, f=&[0 0 0], nil? false
//2025/05/11 17:16:36 g type=[]int, g length=3, g capacity=3, g=[1 3 6], nil? false
//2025/05/11 17:16:36 h type=[]int, h length=0, h capacity=0, h=[], nil? false
//2025/05/11 17:16:36 i type=[]main.champion, i length=2, i capacity=2, i=[{Evelynn [Assassin] [Demon] 3} {Vi [Brawler] [Hextech] 3}], nil? false
//2025/05/11 17:16:36 j type=[3]int, j length=3, j capacity=3, j=[1 3 5]
//2025/05/11 17:16:36 k type=[4]int, k length=4, k capacity=4, k=[1 3 5 7]
