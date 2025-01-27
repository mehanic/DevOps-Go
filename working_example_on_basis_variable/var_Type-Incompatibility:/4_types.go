package main

import "fmt"

type UserId int

func main() {
	idx := 1

	var uid UserId = 42
	fmt.Println("uid = ", uid)

	// даже если базовый тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserId) as int64 in assigement
	// myId := idx
	// fmt.Println(myId)

	myId := UserId(idx)
	fmt.Println("myId = ", myId)

}

// uid =  42
// myId =  1
