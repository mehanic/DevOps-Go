package products

import "github.com/kirigaikabuto/Golang1300Lessons/12/config"

type postgreProductStore struct {
	//element from postgr (CRUD)
}

func NewPostgreProductStore(config config.PostgreConfig) (ProductStore, error) {
	//connection to postgre
	return &postgreProductStore{}, nil
}

func (p *postgreProductStore) Create(product *Product) (*Product, error) {
	return nil, nil
}

func (p *postgreProductStore) List() ([]Product, error) {
	return nil, nil
}

func (p *postgreProductStore) GetById(id string) (*Product, error) {
	return nil, nil
}
