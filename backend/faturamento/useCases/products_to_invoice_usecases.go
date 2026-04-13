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