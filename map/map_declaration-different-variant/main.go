package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}
	fmt.Println(myMap)
	//fmt.Println(myMap[0]) This line tries to access a map element using 0 as the key. However, this will cause a compilation error because 0 is an integer, and the map's keys are of type string.
	fmt.Println(myMap["Ahmet"], myMap["Selim"])
	fmt.Println(myMap["Ahmet Mehmet"])

	fmt.Println("-------------------------------")

	myMap1 := map[string]int{ // LITERAL
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}
	fmt.Println(myMap1)
	fmt.Println(myMap1["Aslı"])

	fmt.Println("-------------------------------")

	// key - value aynı veri tipinde olmasına gerek yok

	isMarried := map[string]bool{
		"Ahmet":  true,
		"Ayşe":   false,
		"Mahmut": false,
	}
	fmt.Println(isMarried)

	fmt.Println("-------------------------------")

	myMap2 := make(map[string]int)
	fmt.Println(myMap2)
	fmt.Println(myMap2["Test"])

	fmt.Println("-------------------------------")
	studentGrades1 := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}
	fmt.Println(studentGrades1)
	fmt.Println(studentGrades1["Arin"])
	fmt.Println(studentGrades1["Çınar"])
	fmt.Println(studentGrades1["Elis"])

	/* 	value, ok := studentGrades["Arin"]
	   	fmt.Println(value, ok) */

	/* 	_, ok := studentGrades["Elis"]
	   	fmt.Println(ok) */

	fmt.Println(studentGrades1)
	studentGrades1["Mahmut"] = 55
	fmt.Println(studentGrades1)
	fmt.Println(len(studentGrades1))
	delete(studentGrades1, "Selim")
	fmt.Println(studentGrades1)
	fmt.Println(len(studentGrades1))

	sg := studentGrades1 // maps --> pass by reference
	fmt.Println(studentGrades1)
	fmt.Println(sg)
	delete(sg, "Arin")
	fmt.Println(sg)
	fmt.Println(studentGrades1)

	fmt.Println("-------------------------------")

	studentGrades := map[string]int{
		"Arin":  80,
		"Ahmet": 29,
		"Selim": 72,
		"Ayşe":  77,
		"Çınar": 0,
	}

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}

	studentGrades["AhmetDeniz"] = 100

	fmt.Println(studentGrades)

	//OLMAYAN KEYLER ICIN ZERO VALUE DONER
	value, ok := studentGrades["Arin"]
	fmt.Println(value, ok)

	//---
	fmt.Println("-------------------------------")

	myArr := [5]string{"tahran", "belgrad", "roma", "tiflis", "moskova"}
	for index, value := range myArr {
		fmt.Println(index, value)
	}
	fmt.Println()
	mySlc := []int{1, 2, 3, 4, 5}
	for i, v := range mySlc {
		fmt.Println(i, v)
	}

	fmt.Println("-------------------------------")

	myMap3 := map[string][]string{
		"Yaşar Kemal":     {"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  {"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": {"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}

	fmt.Println(myMap3)
	fmt.Println(myMap3["Sabahattin Ali"])
	fmt.Println(myMap3["Haruki Murakami"][0])

	for key, value := range myMap3 {
		fmt.Println("Our author: ", key)
		fmt.Println("Some of their books are listed below:")
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}

	fmt.Println("-------------------------------")

	mySlcExample := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	mySlc1 := mySlcExample[:4]
	fmt.Println(mySlc1)
	mySlc2 := mySlcExample[4:7]
	fmt.Println(mySlc2)
	mySlc3 := mySlcExample[6:]
	fmt.Println(mySlc3)
	mySlc4 := append(mySlcExample[2:4], mySlcExample[6:8]...)
	fmt.Println(mySlc4)

	fmt.Println("-------------------------------")

}