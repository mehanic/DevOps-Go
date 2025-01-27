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

	for index := 0; index < 4; index++ {
		grade := grades[index]
		fmt.Printf("my name is %v grade= %.2f Type is %v \n", students[index], grade, calculateGrade(grade))
	}

}

// student [Kutui Munbod Khing Mali]
// grade [1.5 4 2 3]
// my name is Kutui grade= 1.50 Type is D 
// my name is Munbod grade= 4.00 Type is A 
// my name is Khing grade= 2.00 Type is D 
// my name is Mali grade= 3.00 Type is B 
