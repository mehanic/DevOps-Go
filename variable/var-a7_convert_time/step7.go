package main //Scanln and TIME7

import (
	"fmt"
	"time"
)

func main() {

	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello -->", name) // Hello --> Alxan)

	vaqt := time.Date(2024, time.October, 11, 23, 47, 78, 45, time.Local)
	fmt.Println(vaqt)

	fmt.Println(time.Now())

}


// Hello --> Peter
// 2022-02-11 23:48:18.000000045 +0100 CET
// 2024-10-01 23:02:09.437547005 +0200 CEST m=+1.445684591
