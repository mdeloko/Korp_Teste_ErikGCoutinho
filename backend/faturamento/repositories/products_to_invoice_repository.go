package repositories

import (
	"database/sql"
	"fmt"

	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
)

type ProductsToInvoicesRepository struct{
	connection *sql.DB
}

func NewProductsToInvoicesRepository(conn *sql.DB) ProductsToInvoicesRepository {
	return ProductsToInvoicesRepository{
		connection: conn,
	}
}

func (ptir *ProductsToInvoicesRepository) GetProductsToInvoices() ([]models.ProductToInvoice, error){
	query := "SELECT * FROM products_to_invoices;"
	rows,err := ptir.connection.Query(query)

	if err!=nil {
		fmt.Println(err)
		return []models.ProductToInvoice{}, err
	}

	var ptirList []models.ProductToInvoice
	var ptirObj models.ProductToInvoice

	for rows.Next() {
		err = rows.Scan(
			&ptirObj.Id,
			&ptirObj.Invoice_id,
			&ptirObj.Product_id,
			&ptirObj.Amount)
		if err != nil{
			fmt.Println(err)
			return []models.ProductToInvoice{}, err
		}
		ptirList = append(ptirList, ptirObj)
	}
	rows.Close()
	return ptirList, nil
}