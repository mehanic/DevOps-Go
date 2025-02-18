package variables

// This is how we import packages

import "fmt"

func StringCreator() {

	// This is how we create string variables

	var name string = "Berkay"

	// This is how we can also create variables in easy way

	name2 := "Berkay in easy way"

	// This is how we print string variables with fmt package

	fmt.Println(name)
	fmt.Println(name2)

	// If we want to see data type, we can use %T

	fmt.Printf("Data Type : %T\n", name)

}

func IntegerCreator() {

	// This is how we create integer variables

	var age int = 27

	// This is how we can also create variables in easy way

	age2 := 28

	// This is how we print integer variables with fmt package

	fmt.Println("Age : " + fmt.Sprintf("%v", age))
	fmt.Println("Age2 : " + fmt.Sprintf("%v", age2))

	// If we want to see data type, we can use %T

	fmt.Printf("Data Type : %T\n", age)

}

func FloatCreator() {

	// This is how we create float variables

	var price float32 = 100.50

	// This is how we can also create variables in easy way

	price2 := 28.75

	// This is how we print float variables with fmt package

	fmt.Println("Price : " + fmt.Sprintf("%v", price))
	fmt.Println("Price2 : " + fmt.Sprintf("%v", price2))

	// If we want to see data type, we can use %T

	fmt.Printf("Data Type : %T\n", price)

}

func BooleanCreator() {

	// This is how we create boolean variables

	var status bool = false

	// This is how we can also create variables in easy way

	status2 := true

	// This is how we print boolean variables with fmt package

	fmt.Println("Status : " + fmt.Sprintf("%v", status))
	fmt.Println("Status2 : " + fmt.Sprintf("%v", status2))

	// If we want to see data type, we can use %T

	fmt.Printf("Data Type : %T\n", status)

}
