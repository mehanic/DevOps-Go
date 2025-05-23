package main

import "fmt"

func calculateGrade(input_grade float32) string {
	if input_grade < 1.5 {
		return "F"
	} else if input_grade < 3 {
		return "D"
	} else if input_grade < 3 {
		return "C"
	} else if input_grade < 4 {
		return "B"
	} else {
		return "A"
	}
}
func main() {
	calculateGrade(3.2)
	// const name, age = "Kim", 22
	// fmt.Print(name, " is ", age, " years old.")
	const name = "Boonkwang"
	const studentNo = 1

	// fmt.Printf("My name is %v and Student Number is %d\n", name, studentNo)
	fmt.Print("My name is ", name, "\nMy student number is ", studentNo, "\n")

	var teacher1 string = "goodnight"
	teacher2 := "goodnight"
	fmt.Println("teacher1", teacher1)
	fmt.Println("teacher2", teacher2)
	// change name
	teacher1 = name
	fmt.Println("teacher1", teacher1)
	// name = "John"
	fmt.Println("name", name)

	const grade float32 = 3.24324
	fmt.Println("grade=", grade)

	// array [A,B,C,D]
	var students [4]string
	students[0] = "Kutui"
	students[1] = "Munbod"
	students[2] = "Khing"
	students[3] = "Mali"
	fmt.Println("student", students)
	var grades [4]float32
	grades[0] = 1.5
	grades[1] = 4
	grades[2] = 2
	grades[3] = 3
	fmt.Println("grade", grades)

	// fmt.Printf("my name is %v grade= %.2f \n", students[0], grades[0])
	// fmt.Printf("my name is %v grade= %.2f \n", students[1], grades[1])
	// fmt.Printf("my name is %v grade= %.2f \n", students[2], grades[2])
	// fmt.Printf("my name is %v grade= %.2f \n", students[3], grades[3])

	for index := 0; index < 4; index++ {
		// fmt.Println("do  thing")
		fmt.Printf("index %d my name is %v grade= %.2f \n", index, students[index], grades[index])
	}
	fmt.Println("do other thing")

	// homework
	var actors [6]string
	actors[0] = "Khawtu"
	actors[1] = "Gookgik"
	actors[2] = "Boonkwang"
	actors[3] = "Meemee"
	actors[4] = "Nadech"
	actors[5] = "Yaya"

	for index := 0; index < 6; index++ {
		println("Actor name :", actors[index])
	}

	// x := 10
	// y := 0
	// for index := 0; index < 10; index++ {
	// 	fmt.Println(x + y)
	// 	y = y + 10
	// }

	// // [1 * 10,2* 10,3* 10,4* 10,5* 10,6* 10,7* 10,8* 10,9* 10,10* 10]
	// const numberOfLoop = 100 / 10
	// var no int = 1
	// for index := 0; index < numberOfLoop; index++ {
	// 	println(no * 10)
	// 	no = no + 1
	// }

	// for index := 1; index <= numberOfLoop; index++ {
	// 	println("index=", index*10)
	// }

	x := 100
	y := 0
	for i := 0; i < 5; i++ {
		fmt.Println("X = ", x+y)
		y = y + 100
	}

	for index := 1; index <= 10; index++ {
		fmt.Println(index * 2)
	}
	const doingSomething = false
	if doingSomething == true {
		fmt.Println("I'm doing")
	} else {
		fmt.Println("Sorry I'm very lazy")
	}

	var exampleBool bool = true
	if exampleBool == true {
		fmt.Printf("value of example is %v \n", exampleBool)
	}
	ii := 2
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Printf("grade is %v \n", calculateGrade(3.2))
}

// const name, age = "Kim", 22
// fmt.Println(name, "is", age, "years old.")
// const name, age = "Kim", 22
//  fmt.Printf("%s is %d years old.\n", name, age)


//output
// My name is Boonkwang
// My student number is 1
// teacher1 goodnight
// teacher2 goodnight
// teacher1 Boonkwang
// name Boonkwang
// grade= 3.24324
// student [Kutui Munbod Khing Mali]
// grade [1.5 4 2 3]
// index 0 my name is Kutui grade= 1.50 
// index 1 my name is Munbod grade= 4.00 
// index 2 my name is Khing grade= 2.00 
// index 3 my name is Mali grade= 3.00 
// do other thing
// Actor name : Khawtu
// Actor name : Gookgik
// Actor name : Boonkwang
// Actor name : Meemee
// Actor name : Nadech
// Actor name : Yaya
// X =  100
// X =  200
// X =  300
// X =  400
// X =  500
// 2
// 4
// 6
// 8
// 10
// 12
// 14
// 16
// 18
// 20
// Sorry I'm very lazy
// value of example is true 
// two
// grade is B 