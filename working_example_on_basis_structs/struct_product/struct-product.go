package main

import (
	"fmt"
	"time"
)

// Define the Product and Price structures
type Product struct {
	ID    int
	Name  string
	Price *Price
}

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
		{ID: 2, ProductID: 2, Amount: 1.99, Date: time.Date(2015, time.December, 5, 11, 11, 17, 0, time.UTC)},
	},
	3: {
		{ID: 3, ProductID: 3, Amount: 3.00, Date: time.Date(2015, time.December, 20, 11, 11, 17, 0, time.UTC)},
	},
}

// CompleteProduct combines product data with the latest price
func CompleteProduct(p *Product) *Product {
	if p == nil {
		return nil
	}
	productPrices, exists := prices[p.ID]
	if !exists {
		return nil
	}

	// Find the price with the latest date
	var latestPrice *Price
	for _, price := range productPrices {
		if latestPrice == nil || price.Date.After(latestPrice.Date) {
			latestPrice = &price
		}
	}

	// Create a copy of the product and add the latest price
	result := *p
	result.Price = latestPrice
	return &result
}

// GetProduct retrieves a product by ID and completes it with the latest price
func GetProduct(productID int) *Product {
	p, exists := products[productID]
	if !exists {
		return nil
	}
	return CompleteProduct(&p)
}

// ManyProducts retrieves multiple products by IDs and completes them
func ManyProducts(ids []int) []*Product {
	result := make([]*Product, len(ids))
	for i, id := range ids {
		result[i] = GetProduct(id)
	}
	return result
}

// AllProducts retrieves all products and completes them
func AllProducts() []*Product {
	result := make([]*Product, 0, len(products))
	for _, p := range products {
		result = append(result, CompleteProduct(&p))
	}
	return result
}

func main() {
	fmt.Println("printing all...")
	ps := AllProducts()
	for _, p := range ps {
		fmt.Println(p)
	}

	fmt.Println("printing all, but by all ids...")
	ids := make([]int, len(ps))
	for i, p := range ps {
		if p != nil {
			ids[i] = p.ID
		}
	}
	ps = ManyProducts(ids)
	for _, p := range ps {
		fmt.Println(p)
	}
}

// printing all...
// &{3 HTML Book 0xc000090180}
// &{1 Python Book 0xc0000901e0}
// &{2 JavaScript Book 0xc000090240}
// printing all, but by all ids...
// &{3 HTML Book 0xc000090270}
// &{1 Python Book 0xc0000902d0}
// &{2 JavaScript Book 0xc000090330}

