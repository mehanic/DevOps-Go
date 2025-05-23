package main

import "fmt"

func main() {

	fmt.Println("Hello world")
	//uint8 : unsigned 8-bit integers(0 to 255)
	//uint16 : unsigned 16-bit integers(0 to 65535)
	//uint32 : unsigned 32-bit integers(0 to 4294967295)
	//uint64 : unsigned 64-bit integers(0 to 18446744073709)
	//int8 : signed 8-bit integers(-128 to -127)
	//int16 : signed 16-bit integers(-32768 to 32767)
	//int32 : signed 32-bit integers(-2147483648 to 2147483647)
	//int64 : signed 64-bit integers(-9223372036854775808 to 9223372036854775807)

	//var age int = 26
	//var favNum float64 = 1.66233
	randNum := 1
	fmt.Println("random number is", randNum)

	//Arithmatic operations

	fmt.Println("6 + 4 is", 6+4)
	fmt.Println("6 - 4 is", 6-4)
	fmt.Println("6 * 4 is", 6*4)
	fmt.Println("6 / 4 is", 6/4)
	fmt.Println("6 % 4 is", 6%4)

	// Language features
	const pit float64 = 3.14159265
	var myName string = "Rehan Kodekar"
	fmt.Println(len(myName), "\n")
	fmt.Println(myName + " is a robot")

	// boolean datatypes
	var isOver40 bool = true
	fmt.Printf("%f \n", pit)
	fmt.Printf("%.3f \n", pit)
	fmt.Printf("%T \n", pit)
	fmt.Printf("%t \n", isOver40)

	//getting integer
	fmt.Printf("%d \n", 100)

	// binary representations
	fmt.Printf("%b \n", 100)

	// getting the character that has character code of a number
	fmt.Printf("%c \n", 44)

	// print hex code
	fmt.Printf("%x \n", 17)

	// scientific notations
	fmt.Printf("%e \n", pit)

	// logical operators
	fmt.Println("true && false = ", true && false)

	fmt.Println("true || false = ", true || false)

	fmt.Println("!true && false = ", !true)

	//loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
    
    // relations operator
    //== != < > <= >=
    
    for j:= 0; j < 5; j++ {
        fmt.Println(j)
    }
    
    yourAge := 18
    
    if yourAge >= 16 {
        fmt.Println("You can drive")
    } else {
        fmt.Println("You cannot drive")
    }
    
    if yourAge >=18 {
        fmt.Println("You can vote")
    } else if yourAge < 18 {
        fmt.Println("You cannot vote")
    } else {
        fmt.Println("You can have fun")
    }
    
    switch yourAge {
        case 16: fmt.Println("Go drive")
        case 18: fmt.Println("Go vote")
        default: fmt.Println("Have fun")
    }
    
    var favNums2[5] float64
    favNums2[0] = 163
    favNums2[1] = 78557
    favNums2[2] = 682
    favNums2[3] = 2.109
    favNums2[4] = 1.177
    
    fmt.Println(favNums2[3])
    
    favNums3 := [5]float64 {1,2,3,4,5}
    
    for i, value := range favNums3 {
        fmt.Println(value, i)
    }
    
    for _, value := range favNums3 {
        fmt.Println(value)
    }
    
    // slicing
    numSlice := []int {5, 4, 3, 2, 1}
    numSlice2 := numSlice[3:5]
    fmt.Println("numSlice[1] = ", numSlice2[1])
    
    // maps
    presAge := make(map[string]int)
    presAge["Rajendra prasad"] = 60
    fmt.Println(len(presAge))
    
    presAge["Apj abul kalam"] = 62
    fmt.Println(len(presAge))

}


// Hello world
// random number is 1
// 6 + 4 is 10
// 6 - 4 is 2
// 6 * 4 is 24
// 6 / 4 is 1
// 6 % 4 is 2
// 13 

// Rehan Kodekar is a robot
// 3.141593 
// 3.142 
// float64 
// true 
// 100 
// 1100100 
// , 
// 11 
// 3.141593e+00 
// true && false =  false
// true || false =  true
// !true && false =  false
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// 0
// 1
// 2
// 3
// 4
// You can drive
// You can vote
// Go vote
// 2.109
// 1 0
// 2 1
// 3 2
// 4 3
// 5 4
// 1
// 2
// 3
// 4
// 5
// numSlice[1] =  1
// 1
// 2
