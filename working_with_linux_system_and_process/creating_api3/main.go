package main

// We will use gin gonic module here: https://github.com/gin-gonic/gin

import (
	"errors"
	"net/http"

	//"vendor/golang.org/x/net/route"
	"github.com/gin-gonic/gin"
)

type product struct {
	ProductCode  string  `json:"id"`
	ProductName  string  `json:"name"`
	MainCategory string  `json:"main"`
	SubCategory  string  `json:"sub"`
	Price        float64 `json:"price"`
	Brand        string  `json:"brand"`
	Active       bool    `json:"isactive"`
}

var products = []product{
	{ProductCode: "42178322", ProductName: "Lloyd 1.5 Ton 3 Star Inverter Split Ac", MainCategory: "Appliances", SubCategory: "Air Conditioners", Price: 2999.99, Brand: "Llold", Active: true},
	{ProductCode: "62432353", ProductName: "Skybags Brat Black 46 Cms Casual Backpack", MainCategory: "Bags & luggage", SubCategory: "Backpacks", Price: 659.99, Brand: "Skybags", Active: true},
	{ProductCode: "34253562", ProductName: "Glowic Unisex's Backpack for Men Women Boys", MainCategory: "Bags & luggage", SubCategory: "Backpacks", Price: 599.90, Brand: "Glowic", Active: true},
	{ProductCode: "53242635", ProductName: "WOW Skin Science Vitamin C Serum for Skin whitenening", MainCategory: "Beauty & health", SubCategory: "Beauty & Grooming", Price: 219.99, Brand: "WOW", Active: false},
}

func main() {

	router := gin.Default()
	router.GET("/products", GetProducts)
	router.GET("/productdetails/:id", ReturnProductDetails)
	router.POST("/products", CreateProduct)
	router.PATCH("/activateproduct", MakeProductActive)
	router.PATCH("/inactivateproduct", MakeProductPassive)
	//router.PATCH("/filter", ReturnFilteredProducts)
	router.Run("localhost:8080")

}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductDetails(id string) (*product, error) {
	for i, b := range products {
		if b.ProductCode == id {
			return &products[i], nil
		}
	}

	return nil, errors.New("Product Not Found")
}

func ReturnProductDetails(c *gin.Context) {
	id := c.Param("id")
	product, err := GetProductDetails(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var new_product product

	if err := c.BindJSON(&new_product); err != nil {
		return
	}

	products = append(products, new_product)
	c.JSON(http.StatusCreated, new_product)

}

func MakeProductActive(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing ID Query Parameter"})
		return
	}

	product, err := GetProductDetails(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}

	if product.Active {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product Already Active!"})
		return
	}

	product.Active = true
	c.JSON(http.StatusOK, product)
	c.JSON(http.StatusOK, gin.H{"message": "Product activated!"})

}

func MakeProductPassive(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing ID Query Parameter"})
		return
	}

	product, err := GetProductDetails(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}

	if !product.Active {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product Already Inactive!"})
		return
	}

	product.Active = false
	c.JSON(http.StatusOK, product)
	c.JSON(http.StatusOK, gin.H{"message": "Product inactivated!"})

}

/*
func FilterProducts(isactive bool) (*product, error) {
	for i, b := range products {
		if b.Active == isactive {
			return &products[i], nil
		}
	}

	if isactive {
		return nil, errors.New("There is no active product")
	} else {
		return nil, errors.New("There is no inactive product")
	}
}

func ReturnFilteredProducts(c *gin.Context) {
	isactive := c.Param("isactive")
	productlist, err := GetProductDetails(isactive)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}

	c.JSON(http.StatusOK, productlist)
}
*/
