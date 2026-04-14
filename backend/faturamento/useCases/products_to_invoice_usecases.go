package usecases

import (
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
)

type ProductsToInvoiceUseCases struct {
	repository repositories.ProductsToInvoicesRepository
}

func NewProductsToInvoiceUseCase(repo repositories.ProductsToInvoicesRepository) ProductsToInvoiceUseCases {
	return ProductsToInvoiceUseCases{
		repository: repo,
	}
}

func (ptiu *ProductsToInvoiceUseCases) GetProductsToInvoices() ([]models.ProductToInvoice, error) {
	return ptiu.repository.GetProductsToInvoices()
}

func (ptiu *ProductsToInvoiceUseCases) AddProductToInvoice(invoice_id int, product_id string, amount int) error {
	return ptiu.repository.AddProductToInvoice(invoice_id,product_id,amount)
}

func (ptiu *ProductsToInvoiceUseCases) RemoveProductFromInvoice(id int) error {
	return ptiu.repository.RemoveProductFromInvoice(id)
}