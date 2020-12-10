package product

import "github.com/golangRestApi/helper"

type Service interface {
	GetProducts(params *getProductsRequest) (*ProductList, error)
}

type service struct {
	//se define la estructura para acceder a repository
	repo Repository
}

func NewService(repo Repository) Service {
	//se asigna instancia de repository sobre la propiedad repo
	return &service{
		repo: repo,
	}
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)

	helper.Catch(err)
	totalProducts, err := s.repo.GetTotalProducts()

	helper.Catch(err)
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}
