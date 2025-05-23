package main

import "fmt"

func main() {

	//compex data types one of them collections
	//collections -> arrays,map,slices

	cityOne := "istanbul"
	cityTwo := "roma"
	cityThree := "tahran"
	cityFour := "belgrad"
	fmt.Println(cityOne, cityTwo, cityThree, cityFour)

	town := [4]string{"Tokio", "Seul", "Sapporo", "Kioto"}
	fmt.Println(town)

	// cities3 := []string{"istanbul", "roma", "tahran", "belgrad"}
	// fmt.Println(cities3)
	// fmt.Println(cities3[0]) // zero based index
	// fmt.Println(cities3[3])
	// fmt.Println(len(cities3))
	// cities3[3] = "Ankara"
	// fmt.Println(cities3[4]) //staticly length var hata alıyor run time da !!!!!!!

	fmt.Println("----------------------------------")
	/* 	var myArr [5]int
	   	fmt.Println(myArr)
	   	//myArr[0] = "istanbul"
	   	myArr[0] = 100
	   	fmt.Println(myArr)
	   	myArr[len(myArr)-1] = 200
	   	fmt.Println(myArr) */

	// var myArr3 []int //hatta vermedi
	// myArr3[100] = 1  //run time error
	// fmt.Println(myArr3)

	/*var myArr [4]int
	var myArr2 [5]int

	if myArr == myArr2 {
		fmt.Println("MESAJ")
	}*/ //uzunlıkta veri tipine dahil hata veriri

	cities1 := [4]string{"istanbul", "roma", "tahran", "belgrad"}
	for i := 0; i < len(cities1); i++ {
		fmt.Println(i, cities1[i])
	}
	cities1[0] = "ANKARA"
	fmt.Println()
	for i := 0; i < len(cities1); i++ {
		fmt.Println(i, cities1[i])
	}

	// FOR --- RANGE
	// for index, value := range myArr
	fmt.Println("----------------------------------")

	cities := [4]string{"istanbul", "roma", "tahran", "belgrad"}

	/* 	for index, city := range cities {
		fmt.Println(index, city)
	} */

	for _, city := range cities {
		fmt.Println(city)
	}

	fmt.Println("----------------------------------")

	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myArr = mySquare(myArr) // First Class Functions
	fmt.Println(myArr)

}

func mySquare(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
