package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	fmt.Println("Hello World")
	fmt.Printf("Hello World")
	// println prints to a new line, print prints on the same line, and printf can directly print strings

	name := "Deniz"

	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name)

	fmt.Print("My name is:", name, "\n")
	fmt.Println("My name is:", name)
	fmt.Printf("My name is:", name)

	/*
		Output:
		My name is:Deniz
		My name is: Deniz
		My name is: %!(EXTRA string=Deniz)
	*/

	fmt.Printf("My name is: %v", name) // prints the value
	fmt.Printf("My name is: %T", name) // prints the type
	fmt.Printf("My name is: %v %T", name, name)
	// %d shows numbers in base 10, %c for characters, but %v works with everything

	// VISIBILITY --> The concept of scope

	// NAMING
	/*
		Names should be simple and understandable.
		CamelCase is used.
		Abbreviations are written in uppercase (e.g., URL instead of Url).
		Other variable names can also use uppercase, like `myHTTP`.
		For loop indices can use short names like i, j, k.
		Variable names can include Turkish characters, but it's best to avoid them for broader compatibility.
	*/

	// SHADOWING -->

	x := 5

	{
		x := 10
		fmt.Println(x) // Outputs 10
	}

	fmt.Println(x) // Outputs 5

}
