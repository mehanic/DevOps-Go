package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kirigaikabuto/Golang1300Lessons/12/config"
	"github.com/kirigaikabuto/Golang1300Lessons/12/products"
	"log"
	"os"
)

func main() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatal(err)
		return
	}
	url := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DB")
	collection := os.Getenv("MONGO_COLLECTION")
	//productPostgreStore, err := products.NewPostgreProductStore(config.PostgreConfig{})
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	productMongoStore, err := products.NewProductStore(config.MongoConfig{
		Url:        url,
		Database:   database,
		Collection: collection,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	productService := products.NewProductService(productMongoStore)
	if err != nil {
		log.Fatal(err)
		return
	}
	getByIdCmd := &products.GetProductByIdCommand{Id: "102821a6-90b8-42fb-90d4-ef7c4d0b7002"}
	product, err := productService.GetProductById(getByIdCmd)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(product)
	//createCmd := &products.CreateProductCommand{
	//	Name:  "asdsadas",
	//	Price: 123,
	//}
	//newProduct, err := productService.CreateProduct(createCmd)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(newProduct)
	//product, err := productStore.GetById("2c613b4c-79dd-4024-8ad8-3d7b2b2f6566")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(product)
	//	p := &products.Product{
	//		Name:  "product2",
	//		Price: 123,
	//	}
	//	newProduct, err := productStore.Create(p)
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//	fmt.Println(newProduct)
	//	productsObjects, err := productStore.List()
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//	fmt.Println(productsObjects)
}
