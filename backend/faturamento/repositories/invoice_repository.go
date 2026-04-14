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
	query := "SELECT id, invoice_status FROM invoices ORDER BY id"
	rows, err := ir.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Invoice{}, err
	}
	defer rows.Close()
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
	
	return invoiceList, nil
}

func (ir *InvoiceRepository) GetInvoice(id int) (models.Invoice, error) {
	query,err := ir.connection.Prepare("SELECT id, invoice_status FROM invoices WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return models.Invoice{}, err
	}

	var invoice models.Invoice
	err = query.QueryRow(id).Scan(&invoice.ID,&invoice.Status)
	if err != nil {
		fmt.Println(err)
		return models.Invoice{}, err
	}
	defer query.Close()
	
	return invoice, nil
}

func (ir *InvoiceRepository) CreateInvoice() (id int64, err error) {
	query, err := ir.connection.Prepare(`
		INSERT INTO invoices (invoice_status)
		VALUES ('opened')
		RETURNING id
	`)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer query.Close()

	err = query.QueryRow().Scan(&id)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return id,nil
}

func (ir *InvoiceRepository) UpdateInvoiceStatus(id int, newStatus string) (updatedInvoice models.Invoice, err error){
	updatedInvoice = models.Invoice{}

	query, err := ir.connection.Prepare(`
		UPDATE invoices 
		SET invoice_status = $1
		WHERE id = $2
		RETURNING id, invoice_status
	`)
	if err != nil {
		fmt.Println(err)
		return updatedInvoice, err
	}
	defer query.Close()

	err = query.QueryRow(newStatus,id).Scan(&updatedInvoice.ID,&updatedInvoice.Status)
	if err != nil {
		fmt.Println(err)
		return updatedInvoice, err
	}
	return updatedInvoice, nil
}

func (ir *InvoiceRepository) DeleteInvoice(id int) (wasDeleted bool, err error) {
	query, err := ir.connection.Prepare(`
		DELETE FROM invoices 
		WHERE id = $1
		RETURNING id, invoice_status
	`)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer query.Close()
	res, err := query.Exec(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if rowAffected == 0 {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}