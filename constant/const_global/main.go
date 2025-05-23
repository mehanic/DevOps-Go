package main

import (
	"fmt"
)

func main() {

	/* 5
	   3.14
	   "passed"
	   true */

	/* 	r := 3.0
	   	const pi float64 = 3.14
	   	areaOfCircle := pi * (math.Pow(r, 2))
	   	fmt.Println(areaOfCircle) */

	/* 	const x int = 2
	   	const y float32 = 3.4
	   	const z string = "test"
	   	const t bool = false
	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", z, z)
		   fmt.Printf("%T, %v\n", t, t) */

	/* 	const x = 2
	   	//x = 5
	   	// x++
	   	// x = x + 1
		   fmt.Printf("%T, %v\n", x, x) */

	// const ---> compile time
	// var ---> run time

	/*
		const x int
		x = 3
		fmt.Println(x) */
	// hata verir declare ettiğim yerde assign etmek zorundayım çünkü compile time da hesaplanır.

	//GO'da değişkenlerle constantlar aynı şekilde adlandırılır.
	//Constantlara tekrar değer atanamaz
	//declare ettiğin yerde atadığın sürece diğer tüm init yöntemleri kullanılabilir

	//var x1 = math.Pow(3, 4) bu okey
	//const x2 = math.Pow(3, 4) bu olmaz hata verir cünkü pow runtime da çalışır.

	/* 	y := 3
	   	const x = y
	   	fmt.Printf("%T, %v\n", y, y)
		   fmt.Printf("%T, %v\n", x, x)
		   bu da hata verir var runtime da hesaplandığı için*/

	/* 	const x = 1
	   	var y = 3
	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)
	   	fmt.Printf("%T, %v\n", x+y, x+y)
	*/

	const x, y = 3, 5

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

}

// int, 3
// int, 5
