package main

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// створення двох зрізів
	someLanguages := []string{"python", "java", "ada", "lisp", "c++"}
	otherLanguages := []string{"c", "go", "javascript", "php", "swift"}
	languages := append(someLanguages, otherLanguages...) // об’єднання зрізів

	fmt.Println("\nSome programming languages are: ")
	fmt.Println(languages)

	// створюємо title-caser для англійської
	title := cases.Title(language.English)

	// перевертання
	reverseSlice(languages)
	fmt.Println("\n" + title.String(languages[len(languages)-1]) + " is a very popular programming language.")
	fmt.Println(title.String(languages[0]) + " is the programming language used to build iOS and Mac OSX applications.")

	// сортування за спаданням
	sort.Sort(sort.Reverse(sort.StringSlice(languages)))
	fmt.Println("\n" + fmt.Sprint(languages))

	// сортування за зростанням
	sort.Strings(languages)
	fmt.Println(title.String(languages[1]) + " is my favourite programming language but " + strings.ToUpper(languages[len(languages)-3]) + " is my worst.")

	// знову за спаданням
	sort.Sort(sort.Reverse(sort.StringSlice(languages)))
	fmt.Println(title.String(languages[len(languages)-1]) + " is one of the oldest languages around.")

	// вставка, додавання, видалення
	languages = insertAt(languages, "ruby", 4)
	languages = append(languages, "rust")
	languages = removeAt(languages, 2)

	fmt.Printf("\nWe have %d languages in the list: \n\t%v\n", len(languages), languages)
	sort.Strings(languages)
	fmt.Printf("\nIn alphabetical order, the list looks like this: \n\t%v\n", languages)

	// видалення конкретних елементів
	languages = removeItem(languages, "go")
	languages = removeAt(languages, len(languages)-1)

	fmt.Printf("\nNow we have %d languages in the list: \n\t%v\n", len(languages), languages)
}

// reverseSlice reverses the order of the slice.
func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// insertAt inserts an item at a specific position in the slice.
func insertAt(slice []string, item string, index int) []string {
	if index > len(slice) {
		index = len(slice)
	}
	slice = append(slice[:index], append([]string{item}, slice[index:]...)...)
	return slice
}

// removeAt removes an item at a specific position in the slice.
func removeAt(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// removeItem removes the first occurrence of a specific item in the slice.
func removeItem(slice []string, item string) []string {
	for i, v := range slice {
		if v == item {
			return removeAt(slice, i)
		}
	}
	return slice
}

// Some programming languages are:
// [python java ada lisp c++ c go javascript php swift]

// Python is a very popular programming language.
// Swift is the programming language used to build iOS and Mac OSX applications.

// [swift python php lisp javascript java go c++ c ada]
// C is my favourite programming language but PHP is my worst.
// Ada is one of the oldest languages around.

// We have 11 languages in the list:
// 	[swift python lisp ruby javascript java go c++ c ada rust]

// In alphabetical order, the list looks like this:
// 	[ada c c++ go java javascript lisp python ruby rust swift]

// Now we have 9 languages in the list:
// 	[ada c c++ java javascript lisp python ruby rust]
