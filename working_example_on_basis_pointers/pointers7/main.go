package main

import "fmt"

func main() {

	name := "arin"

	fmt.Println(name)
	fmt.Println(&name) // & --> address operator

	
	x := 22

	var y *int = &x
	fmt.Printf("%T, %v\n", y, y)
	z := &name
	fmt.Printf("%T, %v\n", z, z)

	fmt.Println(x)        // 22
	fmt.Println(&x)       // 22 nin adresi &x ---> *int
	fmt.Println(*(&x))    // dereferencing
	fmt.Println(&(*(&x))) // 
	fmt.Println(*(&(*(&x))))
	fmt.Println("**************************************************")

	/* 	x1 := 10
	   	x2 := x1
	   	fmt.Println(x1, x2)
	   	x1 = 5
	   	fmt.Println(x1, x2) */

	/* 	x1 := 10
	   	x2 := &x1
	   	fmt.Println(x1, x2)
	   	fmt.Println(x1, *x2)
	   	*x2 = 3
	   	fmt.Println(x1, *x2)
	   	x3 := &x1
	   	*x3 = 5
	   	fmt.Println(x1, *x2, *x3) */

	// x1 := [4]int{1, 10, 100, 1000}  // array pass by value

	x1 := []int{1, 10, 100, 1000}
	x2 := x1

	fmt.Println(x1, x2)

	x2[0] = 3
	fmt.Println(x2)
	fmt.Println(x1) // slice pass by reference

	xx := 5
	fmt.Println(xx)
	double(&xx)
	fmt.Println(xx)

}

func double(num *int) { // double --> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}

/* package main
import "fmt"
func main() {       // GO pass by value
	x := 5
	fmt.Println(x)
	double(x)
	fmt.Println(x)
}
func double(num int) {
	num *= 2
	fmt.Println(num)
} */

/* package main
import "fmt"
func main() {
	mySlc := []int{1, 10, 100}
	fmt.Println(mySlc)
	double(mySlc)
	fmt.Println(mySlc)
}
func double(slc []int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2
	}
	fmt.Println(slc)
} */

/* package main
import "fmt"
func main() {
	myArr := [3]int{1, 10, 100}
	fmt.Println(myArr)
	double(myArr)
	fmt.Println(myArr)
}
func double(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
} */

//Practice *************************************************************************************************

// 1 -) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
// veri tipi ve değerini '10' yazdırınız.

/* package main
import "fmt"
type myType int
func main() {
	var x myType
	x = 10
	fmt.Printf("%T, %v", x, x)
} */

// 2 -) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında
// bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.

/* package main
import "fmt"
func main() {
	x := 10
	fmt.Println(x)
	y := (&x)
	fmt.Println(y)
	*y = 20
	fmt.Println((*y))
	fmt.Println(x)
} */

// 3 -) Underlying Type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.

/* package main
import "fmt"
type rectangle struct {
	a, b int
}
func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları: ", r.a, " ve ", r.b, " olan dikdörtgen ")
}
func (r rectangle) area() int {
	return r.a * r.b
}
func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}
func main() {
	r1 := rectangle{3, 8}
	r1.display()
	fmt.Println("Alanı: ", r1.area())
	fmt.Println("Çevresi: ", r1.circumference())
} */

// 4-) family aile bireyleri şeklinde veri tipi oluşturalım, underlying struct. Aile bireylerinin hepsini farklı
// şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field age, name, isMarried field.

/*
package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {

	family1 := family{
		name:      "Arin",
		age:       5,
		isMarried: false,
	}

	family2 := family{
		name: "Elis",
		age:  3,
	}

	family3 := family{
		"Gurcan",
		40,
		true,
	}

	var family4 family
	family4.name = "Gamze"
	family4.age = 40
	family4.isMarried = true

	return []family{family1, family2, family3, family4}
}

func main() {
	families := showFamily()
	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}
*/
