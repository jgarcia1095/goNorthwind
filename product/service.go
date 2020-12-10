package product

import "github.com/golangRestApi/helper"

type Service interface {
	GetProducts(params *getProductsRequest) (*ProductList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
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
