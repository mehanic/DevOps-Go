package main

import "fmt"

func main() {

	// if <boolean expression> { code }  x%2 == 0 (false)

	/* 	x := 4
	   	if x%2 == 0 {
	   		fmt.Println(x, "çift sayıdır")
	   	} else {
	   		fmt.Println(x, "tek sayıdır")
		   } */

	// if <boolean expression> { code } else { code }

	/* 	if false {
		fmt.Println("Mesaj Gosterilecek")
	} */

	x := -5

	if x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır")
	} else {
		fmt.Println(x, "tek sayıdır")
	}

	// if <boolean expression> { code } else if <boolean expression> else { code }

	/* 	x := 10
	   	if x := -5; x < 0 {
	   		fmt.Println(x, "negatif sayıdır")
	   	} else if x%2 == 0 {
	   		fmt.Println(x, "çift sayıdır")
	   	} else {
	   		fmt.Println(x, "tek sayıdır")
	   	}
		   fmt.Println(x) */

	if y := -25; y < 0 {
		fmt.Println(y, "negatif sayıdır")
		fmt.Println("Benim Adım Arin")
	} else {
		if y%2 == 0 {
			fmt.Println(y, "çift sayıdır")
		} else {
			fmt.Println(y, "tek sayıdır")
		}
	}

	if 3 < 5 && 4 > 2 {
		fmt.Println("helloooo")
	}
}
