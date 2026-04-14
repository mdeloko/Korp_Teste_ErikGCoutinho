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

func (iu *InvoiceUseCase) GetInvoice(id int) (models.Invoice, error){
	return iu.repository.GetInvoice(id)
}

func (iu *InvoiceUseCase) CreateInvoice() (id int64, err error) {
	return iu.repository.CreateInvoice()
}

func (iu *InvoiceUseCase) UpdateInvoiceStatus(id int,newStatus string) (models.Invoice, error) {
	return iu.repository.UpdateInvoiceStatus(id, newStatus)
}

func (iu *InvoiceUseCase) DeleteInvoice(id int) (wasDeleted bool, err error) {
	return iu.repository.DeleteInvoice(id)
}