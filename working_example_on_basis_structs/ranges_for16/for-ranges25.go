package main

import (
    "fmt"
    "sort"
)

// Product represents a product with various attributes.
type Product struct {
    Description string
    Price       float64
    Sold        float64
    Stock       float64
}

// sell updates the sold count and stock of the product.
func (p *Product) sell() {
    p.Sold++
    p.Stock--
}

// print prints the product details sorted by key.
func (p *Product) print() {
    productMap := map[string]interface{}{
        "description": p.Description,
        "price":       p.Price,
        "sold":        p.Sold,
        "stock":       p.Stock,
    }

    keys := make([]string, 0, len(productMap))
    for k := range productMap {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    for _, k := range keys {
        fmt.Printf("%s: %v, ", k, productMap[k])
    }
    fmt.Println()
}

func main() {
    // Initialize the product
    product := &Product{
        Description: "WCG E-Book",
        Price:       2.75,
        Sold:        1500.0,
        Stock:       5700.0,
    }

    // Print the initial product information
    product.print()

    // Sell the product 100 times
    for i := 0; i < 100; i++ {
        product.sell()
    }

    // Print the updated product information
    product.print()
}

// description: WCG E-Book, price: 2.75, sold: 1500, stock: 5700, 
// description: WCG E-Book, price: 2.75, sold: 1600, stock: 5600, 
