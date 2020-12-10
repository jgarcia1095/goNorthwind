package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/golangRestApi/helper"
)

//creacion de rutas para las acciones
func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	//almacena un puntero de tipo server que recibe request , endpoint y Encode response
	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	//implementar ruta
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	return r
}

//Decodifica el request
func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
