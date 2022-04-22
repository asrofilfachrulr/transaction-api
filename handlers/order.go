package handlers

import (
	"net/http"

	ht "github.com/asrofilfachrulr/transaction-api/http"
	"github.com/asrofilfachrulr/transaction-api/models/sql"
	"github.com/asrofilfachrulr/transaction-api/models/web"
	"github.com/asrofilfachrulr/transaction-api/repository"
	"github.com/asrofilfachrulr/transaction-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register godoc
// @Summary create new order
// @Description  create new order of existing customer with available products and payment methods, all are entered by their IDs
// @Tags Order
// @Param Body body web.PostOrderInput true "Entry new order"
// @Produce json
// @Success 201 {object} web.NormalResp{data=web.PostOrderOutput}
// @Router /order [post]
func PostOrder(ctx *gin.Context) {
	tx := ctx.MustGet("db").(*gorm.DB).Begin()

	if err := tx.Error; err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	var input web.PostOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ht.ErrResp(ctx, err.Error())
		return
	}

	service := services.OrderService{
		OrderRepo:         &repository.OrderRepo{DB: tx},
		OrderItemRepo:     &repository.OrderItemsRepo{DB: tx},
		OrderPaymentRepo:  &repository.OrderPaymentRepo{DB: tx},
		PaymentMethodRepo: &repository.PaymentMethodRepo{DB: tx},
	}

	o := &sql.Order{CustomerAddressID: input.CustomerAddrID}
	oi := &[]sql.OrderItem{}
	opm := &[]sql.OrderPayment{}

	for _, it := range input.ProductIDs {
		*oi = append(*oi, sql.OrderItem{
			ProductID: it,
		})
	}

	for _, it := range input.PaymentMethodIDs {
		*opm = append(*opm, sql.OrderPayment{
			PaymentMethodID: it,
		})
	}

	if err := service.CreateOrder(o, oi, opm); err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()

	ht.CreatedResp(ctx, "order successfully created", web.PostOrderOutput{
		OrderID: o.ID,
	})
}
