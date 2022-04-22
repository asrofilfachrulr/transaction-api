package handlers

import (
	"net/http"

	ht "github.com/asrofilfachrulr/transaction-api/http"
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/repository"
	"github.com/asrofilfachrulr/transaction-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary get all products
// @Description  get all products
// @Tags Product
// @Produce json
// @Success 200 {object} web.NormalResp{data=[]sql.Product{}}
// @Router /products [get]
func GetAllProducts(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	service := services.ProductService{
		Repo: &repository.ProductRepo{DB: db},
	}

	p := &[]sql.Product{}
	if err := service.GetAllProduct(p); err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	ht.CreatedResp(ctx, "get list customers", *p)
}
