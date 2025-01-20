package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollar": 1000,
		"euro":   100,
		"ap":     5,
	}
	fmt.Println(money)
	fmt.Println(money["dollar"])

	money2 := map[string]int{
		"dollar": 1000,
		"euro":   100,
		"ap":     5,
	}
	fmt.Println(money2)

	money2["dollar"] = 5000
	fmt.Println(money2)

	delete(money2, "ap")
	fmt.Println(money2)

	fmt.Println(money2["gg"]) //default for int is 0

	val, status := money2["dollar"]
	fmt.Println(val, status)

	val2, status2 := money2["gg"]
	fmt.Println(val2, status2)

	m1 := make(map[string]int) // キーがstring型、値がint型
	m1["foo"] = 100
	m1["baa"] = 200

	// 値を指定しながら宣言
	m := map[string]int{"foo": 300, "baa": 400}

	fmt.Println(m)
	fmt.Println(len(m)) //値の個数を調べる

	// 要素の削除
	delete(m, "foo")
	fmt.Println(m)

	// キーの存在を調べる
	// vに値、okにboolleanの真偽値を返す
	v, ok := m["baa"]
	fmt.Println(v)
	fmt.Println(ok)
}

// map[ap:5 dollar:1000 euro:100]
// 1000
// map[ap:5 dollar:1000 euro:100]
// map[ap:5 dollar:5000 euro:100]
// map[dollar:5000 euro:100]
// 0
// 5000 true
// 0 false
// map[baa:400 foo:300]
// 2
// map[baa:400]
// 400
// true
