package main

import "fmt"

func Grades(score float32) string {
	if score > 80 {
		return "grade = A"
	}
	if score > 70 {
		return "grade = B"
	}
	if score > 60 {
		return "grade = C"
	}
	if score > 50 {
		return "grade = D"
	}
	return "grade = F"
}

func Students(name string) string {
	return "my name is " + name
}

func main() {
	ScoreA := Grades(84)
	ScoreB := Grades(62)
	ScoreC := Grades(49)
	student1 := Students("Goodnight")
	student2 := Students("Manbod")
	student3 := Students("Dukdik")
	fmt.Println(student1, ScoreA)
	fmt.Println(student2, ScoreB)
	fmt.Println(student3, ScoreC)
}

// my name is Goodnight grade = A
// my name is Manbod grade = C
// my name is Dukdik grade = F

