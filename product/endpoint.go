package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

//almacena las estructuras para el envio de parametros
type getProductsRequest struct {
	Limit  int
	Offset int
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}
