package main

import "fmt"

func main() {
	//d := []int{1, 2, 3, 4}
	//values 1 2 3 4 :int
	//index  0 1 2 3 :int
	di := map[string]int{
		"one": 123,
		"two": 345,
	}
	di["three"] = 323
	fmt.Println(di["two"])
	fmt.Println(di["three"])
	for i, v := range di {
		fmt.Println(i, v)
	}
	fmt.Println("--------------------")
	main1()
}

func main1() {
	d := map[string]int{}
	d["one"] = 1
	d["two"] = 2
	fmt.Println(d)
}

// 345
// 323
// one 123
// two 345
// three 323
