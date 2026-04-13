package usecases

import (
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
)

type InvoiceUseCase struct{
	repository repositories.InvoiceRepository
}

func NewInvoiceUseCase(repo repositories.InvoiceRepository) InvoiceUseCase {
	return InvoiceUseCase{
		repository: repo,
	}
}

func (iu *InvoiceUseCase) GetInvoices() ([]models.Invoice, error){
	return iu.repository.GetInvoices()
}