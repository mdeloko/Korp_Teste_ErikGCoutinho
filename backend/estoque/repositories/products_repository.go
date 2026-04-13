package repositories

import (
	"database/sql"
	"fmt"

	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) (ProductRepository){
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT code, description, amount FROM products;"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{},err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next(){
		err = rows.Scan(
			&productObj.Code,
			&productObj.Description,
			&productObj.Amount)
		if err != nil {
			fmt.Println(err)
			return []models.Product{},err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}