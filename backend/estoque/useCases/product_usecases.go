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

func (pu *ProductUseCase) CreateProduct(product models.Product) error {
	return pu.repository.CreateProduct(product)
}

func (pu *ProductUseCase) DecrementProduct(code string, amountToDecrement int) (newAmount int, err error) {
	return pu.repository.DecrementProduct(code,amountToDecrement)
}

func (pu *ProductUseCase) IncrementProduct(code string, amountToIncrement int) (newAmount int, err error) {
	return pu.repository.IncrementProduct(code,amountToIncrement)
}

func (pu *ProductUseCase) RenameProduct(code string, newName string) error {
	return pu.repository.RenameProduct(code,newName)
}