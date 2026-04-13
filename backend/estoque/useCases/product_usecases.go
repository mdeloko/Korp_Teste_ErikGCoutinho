package usecases

import (
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/repositories"
)

type ProductUseCase struct{
	repository repositories.ProductRepository
}

func NewProductUseCase(repo repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}
func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}