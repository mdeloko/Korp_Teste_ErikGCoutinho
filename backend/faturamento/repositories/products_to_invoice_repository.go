package repositories

import (
	"database/sql"
	"fmt"

	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
)

type ProductsToInvoicesRepository struct {
	connection *sql.DB
}

func NewProductsToInvoicesRepository(conn *sql.DB) ProductsToInvoicesRepository {
	return ProductsToInvoicesRepository{
		connection: conn,
	}
}

func (ptir *ProductsToInvoicesRepository) GetProductsToInvoices() ([]models.ProductToInvoice, error) {
	query := "SELECT * FROM products_to_invoices;"
	rows, err := ptir.connection.Query(query)

	if err != nil {
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
		if err != nil {
			fmt.Println(err)
			return []models.ProductToInvoice{}, err
		}
		ptirList = append(ptirList, ptirObj)
	}
	rows.Close()
	return ptirList, nil
}

func (ptir *ProductsToInvoicesRepository) GetProductsToInvoiceByInvoiceId(id int) ([]models.Product, error) {
	query := `
	SELECT p.code, p.description, SUM(pti.amount)
	FROM products_to_invoices pti
	INNER JOIN products p ON pti.product_id = p.code
	WHERE pti.invoice_id = $1
	GROUP BY p.code, p.description`
	rows, err := ptir.connection.Query(query, id)

	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var productObj models.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Code,
			&productObj.Description,
			&productObj.Amount)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}

func (ptir *ProductsToInvoicesRepository) AddProductToInvoice(invoice_id int, product_id string, amount int) error {
	query, err := ptir.connection.Prepare(`
		INSERT INTO products_to_invoices (invoice_id, product_id, amount)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(invoice_id, product_id, amount)
	return err
}

func (ptir *ProductsToInvoicesRepository) RemoveProductFromInvoice(invoice_id int, product_id string) error {
	query, err := ptir.connection.Prepare(`
		DELETE FROM products_to_invoices 
		WHERE invoice_id = $1 AND product_id = $2
	`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(invoice_id, product_id)
	return err
}
