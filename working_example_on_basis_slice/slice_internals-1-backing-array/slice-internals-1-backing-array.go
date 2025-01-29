package main

import s "github.com/inancgumus/prettyslice"

func main() {
	// ages, first and last2 have the same backing arrays
	ages := []int{35, 15, 25}
	first := ages[0:1]
	last2 := ages[1:3]

	ages[0] = 55
	ages[1] = 10
	ages[2] = 20

	// grades and ages have separate backing arrays
	grades := []int{70, 99}
	grades[0] = 50

	s.Show("ages", ages)
	s.Show("ages[0:1]", first)
	s.Show("ages[1:3]", last2)
	s.Show("grades", grades)

	// let's create a new scope
	// 'cause I'm going to use variables with the same name
	{
		// ages and agesArray have the same backing arrays
		agesArray := [3]int{35, 15, 25}
		ages := agesArray[0:3]

		ages[0] = 100
		ages[2] = 50

		s.Show("agesArray", agesArray[:])
		s.Show("agesArray's ages", ages)
	}
}

//----

func main1() {
	// #1: arrays and non-empty slice literals create an array.
	// For the arrays, it's explicit, but for the slices,
	// it's done implicitly, behind the scenes.

	grades := [...]float64{40, 10, 20, 50, 60, 70} // #1
	// grades := []float64{40, 10, 20, 50, 60, 70} // #4

	// #5: let's break the connection
	// #6: comment-out
	// var newGrades []float64
	// newGrades = append(newGrades, grades...)

	// #6: shortcut: []float64(nil) is a nil float64 slice
	// newGrades := append([]float64(nil), grades...)

	// #2: cheap: slicing doesn't allocate new memory (array).
	// front := grades[:3]

	// front := newGrades[:3] // #5

	// #3: sort its first segment
	// sort.Float64s(front)

	// #7: new slices look at the same backing array
	// front, front2, front3, newGrades, they all have the same backing array
	// front2 := front[:3]
	// front3 := front

	s.PrintBacking = true       // #1
	s.MaxPerLine = 7            // #1
	s.Show("grades", grades[:]) // #1
	// s.Show("newGrades", newGrades) // #5
	// s.Show("front", front)         // #2
	// s.Show("front2", front2)       // #7
	// s.Show("front3", front3)       // #7
}
