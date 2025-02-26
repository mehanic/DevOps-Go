package main

import "fmt"


func main() {
	runApplication(true)
}

func endApplication() {
	message := recover()
	fmt.Println("Pesan Error : ", message)
	fmt.Println("Program Berakhir")
}

func runApplication(error bool) {
	defer endApplication()
	// error == true
	if error {
		panic("Terjadi Kesalahan Pada Program")
	}
	fmt.Println("Aplikasi berjalan dengan baik")
}