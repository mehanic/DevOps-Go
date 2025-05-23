package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers := Person{1, "Hoasker"}
	pers.SetName("Hoasker RealZealot")
	fmt.Printf("update person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "RZ",
		Person: Person{
			Id:   2,
			Name: "RealZealot",
		},
	}

	acc.SetName("hoasker.rz")
	fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)

}

// update person: main.Person{Id:1, Name:"Hoasker RealZealot"}
// main.Account{Id:1, Name:"hoasker.rz", Person:main.Person{Id:2, Name:"RealZealot"}}
// 3 [1 2 5]
