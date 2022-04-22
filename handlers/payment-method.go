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
// @Summary get all avalaible payment methods
// @Description  get all avalaible payment methods
// @Tags Payment
// @Produce json
// @Success 200 {object} web.NormalResp{data=[]sql.PaymentMethod{}}
// @Router /payment/methods [get]
func GetAllPaymentMethods(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	service := services.PaymentMethodService{
		Repo: &repository.PaymentMethodRepo{DB: db},
	}

	pm := &[]sql.PaymentMethod{}
	if err := service.GetAllMethod(pm); err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	ht.CreatedResp(ctx, "get list customers", *pm)
}
