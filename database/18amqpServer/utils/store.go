package utils

func GetNumber() string {
	return "3"
}

type Product struct {
	Id    int
	Name  string
	price float64
}

func GetProducts() []Product {
	products := []Product{
		{
			Id:    1,
			Name:  "product1",
			price: 123,
		},
		{
			Id:    2,
			Name:  "product2",
			price: 456,
		},
	}
	return products
}
