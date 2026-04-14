package repositories

import (
	"database/sql"
	"errors"
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
	defer rows.Close()

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
	return productList, nil
}

func (pr *ProductRepository) GetProductById(id int) (models.Product, error){
	query, err := pr.connection.Prepare("SELECT code, description, amount FROM products WHERE id=$1")
	if err != nil{
		fmt.Println(err)
		return models.Product{}, err
	}
	defer query.Close()

	var product models.Product

	err = query.QueryRow(id).Scan(
		&product.Code,
		&product.Description,
		&product.Amount)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, nil
		}
		return models.Product{}, err
	}
	return product, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) error {
	query, err := pr.connection.Prepare("INSERT INTO products (code,description,amount) VALUES ($1,$2,$3)")
	if err != nil{
		fmt.Println(err)
		return err
	}
	defer query.Close()

	_, err = query.Exec(product.Code,product.Description,product.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) DecrementProduct(code string, amountToDecrement int) (newAmount int, err error) {
	query, err := pr.connection.Prepare(`UPDATE products
		SET amount = amount - $1
		WHERE code = $2 AND amount >= $1
		RETURNING code, description, amount`)
	if err != nil {
		return -1,err
	}
	defer query.Close()

	var product models.Product
	err = query.QueryRow(amountToDecrement,code).Scan(
		&product.Code,
		&product.Description,
		&product.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("Estoque insuficiente ou Produto não encontrado.")
		}
		return -1,err
	}
	return product.Amount,nil
}

func (pr *ProductRepository) IncrementProduct(code string, amountToIncrement int) (newAmount int, err error) {
	query, err := pr.connection.Prepare(`UPDATE products
		SET amount = amount + $1
		WHERE code = $2
		RETURNING code, description, amount`)
	if err != nil {
		return -1, err
	}
	defer query.Close()

	var product models.Product
	err = query.QueryRow(amountToIncrement,code).Scan(
		&product.Code,
		&product.Description,
		&product.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("Estoque insuficiente ou Produto não encontrado.")
		}
		return -1, err
	}
	return product.Amount, nil
}

func (pr *ProductRepository) RenameProduct(code string, newName string) error {
	query, err := pr.connection.Prepare(`UPDATE products
		SET description = $1
		WHERE code = $2`)
	if err != nil {
		return err
	}
	defer query.Close()

	res, err := query.Exec(newName,code)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil{
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Produto não encontrado.")
	}

	return nil
}