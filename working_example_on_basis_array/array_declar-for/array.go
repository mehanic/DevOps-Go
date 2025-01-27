package main

import "fmt"

func Array() {
	var sayilar [5]int
	sayilar[3] = 57
	fmt.Println(sayilar)
}

func Array2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İzmir"
	sehirler[2] = "Bursa"
	sehirler[3] = "İstanbul"
	sehirler[4] = "Mersin"

	fmt.Println(sehirler)
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}

func main() {
	Array()
	Array2()

}


// [0 0 0 57 0]
// [Ankara İzmir Bursa İstanbul Mersin]
// Ankara
// İzmir
// Bursa
// İstanbul
// Mersin
