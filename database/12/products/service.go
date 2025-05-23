package products

import (
	"errors"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(cmd *CreateProductCommand) (*Product, error)
	GetProductById(cmd *GetProductByIdCommand) (*Product, error)
}

type productService struct {
	//connection to database
	database ProductStore
}

func NewProductService(db ProductStore) ProductService {
	return &productService{database: db}
}

func (p *productService) CreateProduct(cmd *CreateProductCommand) (*Product, error) {
	if cmd.Name == "" {
		return nil, errors.New("please write the name of product")
	}
	if cmd.Price == 0 {
		return nil, errors.New("please write the price of product")
	}
	product := &Product{
		Name:  cmd.Name,
		Price: cmd.Price,
	}
	id := uuid.New()
	product.Id = id.String()
	newProduct, err := p.database.Create(product)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (p *productService) GetProductById(cmd *GetProductByIdCommand) (*Product, error) {
	if cmd.Id == "" {
		return nil, errors.New("please write the id of product")
	}
	product, err := p.database.GetById(cmd.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
