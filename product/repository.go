package product

import (
	"database/sql"

	"github.com/golangRestApi/helper"
)

//Interface para agregar comportamiento adicional a una estructura
type Repository interface {
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
}

type repository struct {
	//se agrega propiedad db para almacenar la conexion
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	//Retorna repository con la conexion
	return &repository{db: databaseConnection}
}

//
func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name ,COALESCE(description,''),standard_cost ,list_price ,category 
	FROM products
	LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	helper.Catch(err)

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products, err
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = `SELECT COUNT(*) FROM products`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}
