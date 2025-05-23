package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserData(promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}

	return cleanedInput
}


func generateSimpleId() string {
	return strconv.Itoa(rand.Intn((1000-10) + 10))
}

type Product struct {
	ID,
	Title,
	Description, Price string
}

func createProduct(title, desc, price string) *Product {
	return &Product{generateSimpleId(), title, desc, price}
}

func (product *Product) getProductInfo() {
	fmt.Printf("ID: %v\nTitle: %v\nDesc: %v\nPrice: %v$\n", product.ID, product.Title, product.Description, product.Price)
}

func Init() {
	title := getUserData("Enter Product name: ")
	desc := getUserData("Shortly describe your product: ")
	price := getUserData("Enter Product price in USD: ")

	createProduct(title, desc, price).getProductInfo()

}


func main() {
	Init()
}


// Enter Product name: car
// Shortly describe your product: for travel
// Enter Product price in USD: 3000
// ID: 583
// Title: car
// Desc: for travel
// Price: 3000$
