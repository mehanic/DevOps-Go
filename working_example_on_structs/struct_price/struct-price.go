package main

import (
	"fmt"
	"time"
)

// Product structure
type Product struct {
	ID    int
	Name  string
	Price float64 // Added this field to hold the latest price
}

// Price structure
type Price struct {
	ID        int
	ProductID int
	Amount    float64
	Date      time.Time
}

// Sample data
var products = map[int]Product{
	1: {ID: 1, Name: "Python Book"},
	2: {ID: 2, Name: "JavaScript Book"},
	3: {ID: 3, Name: "HTML Book"},
}

var prices = map[int][]Price{
	1: {
		{ID: 1, ProductID: 1, Amount: 2.75, Date: time.Date(2016, time.February, 1, 11, 11, 17, 0, time.UTC)},
		{ID: 4, ProductID: 1, Amount: 3.99, Date: time.Date(2016, time.February, 5, 11, 11, 17, 0, time.UTC)},
	},
	2: {
		{ID: 2, ProductID: 2, Amount: 1.10, Date: time.Date(2015, time.January, 5, 11, 11, 17, 0, time.UTC)},
		{ID: 5, ProductID: 2, Amount: 1.99, Date: time.Date(2015, time.December, 5, 11, 11, 17, 0, time.UTC)},
	},
	3: {
		{ID: 3, ProductID: 3, Amount: 3.00, Date: time.Date(2015, time.December, 20, 11, 11, 17, 0, time.UTC)},
	},
}

// CompleteProduct adds the latest price to a product
func CompleteProduct(p *Product) *Product {
	if p == nil {
		return nil
	}
	productPrices, exists := prices[p.ID]
	if !exists {
		return nil
	}

	// Find the price with the latest date
	var latestPrice Price
	for _, price := range productPrices {
		if price.Date.After(latestPrice.Date) {
			latestPrice = price
		}
	}

	// Add the latest price to the product
	p.Price = latestPrice.Amount
	return p
}

// GetProduct retrieves a product by ID and adds the latest price
func GetProduct(productID int) *Product {
	p, exists := products[productID]
	if !exists {
		return nil
	}
	return CompleteProduct(&p)
}

func main() {
	// Test the code
	product := GetProduct(1)
	if product != nil {
		fmt.Printf("Product: %s, Latest Price: $%.2f\n", product.Name, product.Price)
	} else {
		fmt.Println("Product not found")
	}
}


//Product: Python Book, Latest Price: $3.99