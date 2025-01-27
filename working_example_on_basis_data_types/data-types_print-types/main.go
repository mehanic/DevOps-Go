package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	fmt.Println("Hello World")
	fmt.Printf("Hello World")
	//println yeni satıra geçerek basar, print aynı satıra basar printf'e de string verirsek direkt basar

	name := "Deniz"

	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name)

	fmt.Print("Benim adım:", name, "\n")
	fmt.Println("Benim adım:", name)
	fmt.Printf("Benim adım:", name)

	/*
		Benim adım:Deniz
		Benim adım: Deniz
		Benim adım: %!(EXTRA string=Deniz)
	*/

	fmt.Printf("Benim adım: %v", name) // değeri bastırır
	fmt.Printf("Benim adım: %T", name) // tipi yazdırır
	fmt.Printf("Benim adım: %v %T", name, name)
	// %d 10 luk tabanda gösterir %c char ama %v herşeyle kullanılır

	//VISIBILITY --> scope mantığı

	//NAMING
	/*
		Sade ve anlaşılır olmalı
		camelCase kullanılır.
		Kısatmalar büyük harfle URL gibi Url değil
		Başka değişkenadında da büyük myHTTP gibi
		for dögüsü indexi için i,j,k olabilir
		değişken isminde türkçe karakter olabilir
	*/

	//shadowing -->

	x := 5

	{
		x := 10
		fmt.Println(x) //10
	}

	fmt.Println(x) //5

}

// ┌─────(mehanic)─────(~/structure/data_types)
// └> $ go run /home/mehanic/structure/data_types/05_print_types/main.go
// Hello WorldHello World
// Hello WorldDenizDeniz
// DenizBenim adım:Deniz
// Benim adım: Deniz
// Benim adım:%!(EXTRA string=Deniz)Benim adım: DenizBenim adım: stringBenim adım: Deniz string10
// 5
