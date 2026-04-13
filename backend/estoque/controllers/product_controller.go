package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/mdeloko/Korp_Teste_ErikGCoutinho/useCases"
)

type ProductController struct{
	productUseCase usecases.ProductUseCase
}

func NewProductController(useCase usecases.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: useCase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK,products)
}