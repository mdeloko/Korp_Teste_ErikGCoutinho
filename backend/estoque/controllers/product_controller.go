package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/mdeloko/Korp_Teste_ErikGCoutinho/models"
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

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Malformed payload or missing information!",
		})
		return
	}

	err = pc.productUseCase.CreateProduct(product)
	var pqErr *pq.Error

	if err != nil {
		if errors.As(err, &pqErr) {
            // https://www.postgresql.org/docs/current/errcodes-appendix.html
            switch pqErr.Code {
            case "23505": // unique_violation (Violação de Chave Primária ou UNIQUE constraint)
                ctx.JSON(http.StatusConflict, gin.H{"error": "This product already exists! (duplicated code)."})
                return
            case "23502": // not_null_violation (Faltando algum campo obrigatório)
                ctx.JSON(http.StatusBadRequest, gin.H{"error": "A required field was not filled correctly."})
                return
            case "22001": // string_data_right_truncation (Texto muito longo, como VARCHAR(10))
                ctx.JSON(http.StatusBadRequest, gin.H{"error": "A field character limit was exceeded."})
                return
            }
        }
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":"Internal Server Error in the Product Creation."})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (pc *ProductController) DecrementOrIncrementProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"You must search for an actual product id!",
		})
		return
	}
	var payload struct{
		Amount int `json:"amount" binding:"required"`
	}
	err := ctx.BindJSON(&payload)
	if err !=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid payload or missing amount!",
        })
        return
	}
	var newAmount int

	if payload.Amount < 0 {
		newAmount,err = pc.productUseCase.DecrementProduct(id,-payload.Amount)
	}else {
		newAmount,err = pc.productUseCase.IncrementProduct(id, payload.Amount)
	}

	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
	ctx.JSON(http.StatusOK, gin.H{
        "message":    "Updated the inventory with success!",
        "product_id": id,
        "new_amount": newAmount,
    })
}

func (pc *ProductController) RenameProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"You must search for an actual product id!",
		})
		return
	}
	var payload struct{
		NewName string `json:"newName" binding:"required"`
	}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing new name!",
        })
        return
	}

	err = pc.productUseCase.RenameProduct(id,payload.NewName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
	}
	ctx.JSON(http.StatusOK, gin.H{
        "message":    "Updated the inventory with success!",
        "product_id": id,
        "new_name": payload.NewName,
    })
}