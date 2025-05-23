package products

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductStore interface {
	Create(product *Product) (*Product, error)
	List() ([]Product, error)
	GetById(id string) (*Product, error)
}
