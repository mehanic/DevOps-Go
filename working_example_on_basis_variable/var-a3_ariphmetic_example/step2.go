package main

import "fmt"

func main() {

	a := 10
	b := 20

	total := a + b
	fmt.Println(total)

	modul := a % b
	fmt.Println(modul) // --> 10 // to'g'ri kichik  soning o'zi o'tadi

	num1 := 63
	num2 := 10

	module := num1 % num2
	fmt.Println(module)
	//  _______________________________________________________________________________________________________

	num := 0
	fmt.Println(num)
	num += 10 // num = num + 10 --> 10
	fmt.Println(num)
	num -= 10 // num = num - 10 --> 0
	fmt.Println(num)
	num *= 10 // num = num / 10 --> 0
	num /= 10 // num = num /10 -->
	fmt.Println(num)
	// _______________________________________________________________________________________________________

	box := 10
	box -= 5 // box = 10 - 5 --> 5
	box += 4 // box = 5 + 4 --> 9
	box *= 2 // box = 9 * 2 -->
	box /= 3 // box = 18 / 3 --> 6
	fmt.Println(box)
	//___________________________________________________________________________________________________________

	number := 10
	number++            // number += I bu ham  bo'ladi
	fmt.Println(number) // --> 10

	number--            //  number -= 1 bu ham bo'ladi
	fmt.Println(number) // --> 10
	// _______________________________________________________________________________________________________

	number += 5
	fmt.Println(number) // -->
	number -= 3
	fmt.Println(number) // --> 12

	number /= 2
	fmt.Println(number) // --> 6

	number *= 3
	fmt.Println(number) // --> 18

	// fmt.Printf("        %                  \n",m ,m ,m ,m ,m ,m ,m  ,m)
	// fmt.Printf("         %s %s               \ n,  , m,m,m,         ")
	// fmt.Printf("         %s %s %s %s        \  n,  ,  m,m,m            ")
	// fmt.Printf("         %s %s %s %s %s        \ m  , m,m              ")
	// fmt.Printf("         % % % % % %       \m   ,m,m              ")
	// fmt.Printf("         % % % % % % %    \ m  , m,m              ")
	// fmt.Printf("         %s %s %s % % % % %  \m   , m             ")

}

// 30
// 10
// 3
// 0
// 10
// 0
// 0
// 6
// 11
// 10
// 15
// 12
// 6
// 18
