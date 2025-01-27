package main

import "fmt"

var globalVariable = "Global Variable"            //correct declaration
var globalVariable_2 string = "Global Varibale 2" //correct declaration
//globalVariable := "globale variable 3" // incorrect

func main() {
	fmt.Println("Variables")

	var username string = "Rohan"
	fmt.Println("user name is", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isusername bool = true
	fmt.Println("isusername is", isusername)
	fmt.Printf("Variable is of type: %T \n", isusername)

	var smallfl float32 = 3.141114141
	fmt.Println("smallfl is", smallfl)
	fmt.Printf("Variable is of type: %T \n", smallfl)

	test1 := 28.282828282828
	fmt.Println("test1 is", test1)
	fmt.Printf("Variable is of type: %T \n", test1)

}

// Variables
// user name is Rohan
// Variable is of type: string
// isusername is true
// Variable is of type: bool
// smallfl is 3.1411142
// Variable is of type: float32
// test1 is 28.282828282828
// Variable is of type: float64
