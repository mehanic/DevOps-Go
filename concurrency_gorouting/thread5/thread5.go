package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	// "time"
)

// Repository struct with a map to hold data and a mutex for locking.
type Repository struct {
	repo map[int]map[string]interface{}
	lock sync.Mutex
}

// Create a new entry in the repository.
func (r *Repository) create(entry map[string]interface{}) {
	log.Println("[INFO] waiting for lock")
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Println("[INFO] acquired lock")
	newID := len(r.repo)
	entry["id"] = newID
	r.repo[newID] = entry
	log.Println("[INFO] releasing lock")
}

// Find an entry by ID.
func (r *Repository) find(entryID int) (map[string]interface{}, bool) {
	log.Println("[INFO] waiting for lock")
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Println("[INFO] acquired lock")
	entry, exists := r.repo[entryID]
	log.Println("[INFO] releasing lock")
	return entry, exists
}

// Return all entries.
func (r *Repository) all() map[int]map[string]interface{} {
	log.Println("[INFO] waiting for lock")
	r.lock.Lock()
	defer r.lock.Unlock()

	log.Println("[INFO] acquired lock")
	log.Println("[INFO] releasing lock")
	return r.repo
}

// ProductRepository inherits from Repository and adds products.
type ProductRepository struct {
	Repository
}

// Add a product with description and price.
func (pr *ProductRepository) addProduct(description string, price float64) {
	pr.create(map[string]interface{}{"description": description, "price": price})
}

// PurchaseRepository inherits from Repository and adds purchase-related logic.
type PurchaseRepository struct {
	Repository
	productRepository *ProductRepository
}

// Add a purchase for a product by its ID and quantity.
func (pr *PurchaseRepository) addPurchase(productID, qty int) {
	product, exists := pr.productRepository.find(productID)
	if exists {
		totalAmount := product["price"].(float64) * float64(qty)
		pr.create(map[string]interface{}{"product_id": productID, "qty": qty, "total_amount": totalAmount})
	}
}

// Get total sales for a product by its ID.
func (pr *PurchaseRepository) salesByProduct(productID int) map[string]interface{} {
	sales := map[string]interface{}{"product_id": productID, "qty": 0, "total_amount": 0.0}
	allPurchases := pr.all()
	for _, purchase := range allPurchases {
		if purchase["product_id"].(int) == sales["product_id"].(int) {
			sales["qty"] = sales["qty"].(int) + purchase["qty"].(int)
			sales["total_amount"] = sales["total_amount"].(float64) + purchase["total_amount"].(float64)
		}
	}
	return sales
}

// Buyer simulates a buyer making random purchases.
type Buyer struct {
	name               string
	productRepository  *ProductRepository
	purchaseRepository *PurchaseRepository
}

// Run method for a buyer to randomly purchase products.
func (b *Buyer) run(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		maxProductID := len(b.productRepository.all())
		productID := rand.Intn(maxProductID)
		qty := rand.Intn(100)
		b.purchaseRepository.addPurchase(productID, qty)
	}
}

// ProviderAuditor audits the purchases of a product.
type ProviderAuditor struct {
	productID          int
	purchaseRepository *PurchaseRepository
}

// Run method for the auditor to log sales for a product.
func (pa *ProviderAuditor) run(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("[Auditor] %v\n", pa.purchaseRepository.salesByProduct(pa.productID))
}

// Main function to run the simulation.
func main() {
	log.SetFlags(0) // Set logging format without timestamp for simplicity

	productRepository := &ProductRepository{Repository: Repository{repo: make(map[int]map[string]interface{})}}
	purchaseRepository := &PurchaseRepository{Repository: Repository{repo: make(map[int]map[string]interface{})}, productRepository: productRepository}

	// Add products to repository
	var inputAnotherProduct = true
	for inputAnotherProduct {
		var description string
		var price float64
		fmt.Print("product description: ")
		fmt.Scanln(&description)
		fmt.Print("product price: ")
		fmt.Scanln(&price)
		productRepository.addProduct(description, price)

		var input string
		fmt.Print("continue (y/N): ")
		fmt.Scanln(&input)
		inputAnotherProduct = input == "y"
	}

	// Create buyers and start them in goroutines
	buyers := []*Buyer{
		{name: "carlos", productRepository: productRepository, purchaseRepository: purchaseRepository},
		{name: "juan", productRepository: productRepository, purchaseRepository: purchaseRepository},
		{name: "mike", productRepository: productRepository, purchaseRepository: purchaseRepository},
		{name: "sarah", productRepository: productRepository, purchaseRepository: purchaseRepository},
	}

	var wg sync.WaitGroup
	for _, buyer := range buyers {
		wg.Add(1)
		go buyer.run(&wg)
	}

	wg.Wait()

	// Start auditors for each product
	for productID := range productRepository.all() {
		wg.Add(1)
		auditor := &ProviderAuditor{productID: productID, purchaseRepository: purchaseRepository}
		go auditor.run(&wg)
	}

	wg.Wait()
}

// product description: yui
// product price: 690
// [INFO] waiting for lock
// [INFO] acquired lock
// [INFO] releasing lock
// continue (y/N):
