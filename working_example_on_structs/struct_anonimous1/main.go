package main

import "fmt"

func main() {

	// Embedding structs as fields in another struct
	type Contact struct {
		email, address string
		phone          int
	}

	type Workers struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Workers{
		name:   "Shallom Micah Bawa",
		salary: 4000,
		contactInfo: Contact{
			email:   "micahshallom@gmail.com",
			address: "Kaduna, Nigeria",
			phone:   340008000,
		},
	}

	// Printing the worker information
	fmt.Printf("%#v\n", john)

	// Anonymous Structs
	shallom := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Shallom",
		lastName:  "Micah Bawa",
		age:       23,
	}
	fmt.Printf("%#v\n", shallom)

	// Anonymous Fields
	type Book struct {
		string
		float64
		bool
	}

	// Creating an instance of Book
	b1 := Book{"1984 by George", 10.5, false}
	fmt.Printf("%#v\n", b1)
	fmt.Printf("%#v\n", b1.string)

	// Mixing named fields and anonymous fields
	type Employee struct {
		name   string
		salary int
		bool
	}

	// Creating an Employee instance
	e := Employee{"Shallom", 40000, false}
	fmt.Printf("%#v\n", e)

	// #############################################
	// #############################################

	// Struct with named fields
	type BookDetails struct {
		title  string
		author string
		year   int
	}

	// Creating an instance of BookDetails
	lastBook := BookDetails{title: "Anna Karenina"}
	fmt.Println(lastBook.title)
	fmt.Printf("%#v\n", lastBook)

	// Updating fields
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	// Comparing two structs
	aBook := BookDetails{title: "Anna Karenina", author: "", year: 0}
	fmt.Println(lastBook == aBook) // Compare structs directly

	// Making a copy of a struct
	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)

	// #############################################
	// #############################################

	// Create an instance of book using positional fields
	myNewBook := BookDetails{"The Divine Comedy", "Dante Alighieri", 1320}
	fmt.Println(myNewBook)

	// Create an instance of book using named fields
	bestBook := BookDetails{title: "Animal Farm", year: 1945, author: "George Orwell"}
	fmt.Println(bestBook)

	// Create an instance of book with only title set, leaving others as zero values
	justBook := BookDetails{title: "Just a random book"}
	fmt.Printf("%#v\n", justBook)
	fmt.Println(justBook)
}

// main.Workers{name:"Shallom Micah Bawa", salary:4000, contactInfo:main.Contact{email:"micahshallom@gmail.com", address:"Kaduna, Nigeria", phone:340008000}}
// struct { firstName string; lastName string; age int }{firstName:"Shallom", lastName:"Micah Bawa", age:23}
// main.Book{string:"1984 by George", float64:10.5, bool:false}
// "1984 by George"
// main.Employee{name:"Shallom", salary:40000, bool:false}
// Anna Karenina
// main.BookDetails{title:"Anna Karenina", author:"", year:0}
// {title:Anna Karenina author:Leo Tolstoy year:1878}
// false
// {Anna Karenina  2020} {Anna Karenina  0}
// {The Divine Comedy Dante Alighieri 1320}
// {Animal Farm George Orwell 1945}
// main.BookDetails{title:"Just a random book", author:"", year:0}
// {Just a random book  0}
