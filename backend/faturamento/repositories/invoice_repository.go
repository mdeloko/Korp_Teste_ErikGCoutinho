package repositories

import (
	"database/sql"
	"fmt"

	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
)

type InvoiceRepository struct{
	connection *sql.DB
}

func NewInvoiceRepository(connection *sql.DB) (InvoiceRepository) {
	return InvoiceRepository{
		connection: connection,
	}
}

func (ir *InvoiceRepository) GetInvoices() ([]models.Invoice, error) {
	query := "SELECT id, invoice_status FROM invoices;"
	rows, err := ir.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Invoice{}, err
	}
	var invoiceList []models.Invoice
	var invoiceObj models.Invoice

	for rows.Next(){
		err = rows.Scan(
			&invoiceObj.ID,
			&invoiceObj.Status)
		if err != nil {
			fmt.Println(err)
			return []models.Invoice{}, err
		}
		invoiceList = append(invoiceList, invoiceObj)
	}
	rows.Close()
	return invoiceList, nil
}