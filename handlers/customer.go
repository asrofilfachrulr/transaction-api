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
// @Summary add new customer
// @Description  add new customer
// @Tags Customer
// @Param Body body web.RegisterCustomerInput true "Register new customer"
// @Produce json
// @Success 201 {object} web.NormalResp{data=web.RegisterCustomerOutput}
// @Router /customer [post]
func PostCustomer(ctx *gin.Context) {
	tx := ctx.MustGet("db").(*gorm.DB).Begin()
	if err := tx.Error; err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	service := services.CustomerService{
		CustRepo:     &repository.CustRepo{DB: tx},
		CustAddrRepo: &repository.CustAddrRepo{DB: tx},
	}

	var input web.RegisterCustomerInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ht.ErrResp(ctx, err.Error())
		return
	}

	c := &sql.Customer{
		CustomerName: input.Name,
	}
	ca := &sql.CustomerAddress{
		Address: input.Address,
	}

	if err := service.CreateCustomer(c, ca); err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()

	ht.CreatedResp(ctx, "customer created", web.RegisterCustomerOutput{
		CustomerID:     c.ID,
		CustomerAddrID: ca.ID,
		Name:           c.CustomerName,
		Address:        ca.Address,
	})
}

// Register godoc
// @Summary get all customers
// @Description  get all customers
// @Tags Customer
// @Produce json
// @Success 200 {object} web.NormalResp{data=[]web.CustomerOutput{}}
// @Router /customers [get]
func GetAllCustomers(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	service := services.CustomerService{
		CustRepo: &repository.CustRepo{DB: db},
	}

	cWeb := []web.CustomerOutput{}
	cSql := &[]sql.Customer{}
	if err := service.GetAllCustomer(cSql); err != nil {
		ht.ErrResp(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, c := range *cSql {
		for _, ca := range c.CustomerAddresses {
			cWeb = append(cWeb, web.CustomerOutput{
				CustomerID:     c.ID,
				CustomerAddrID: ca.ID,
				Name:           c.CustomerName,
			})
		}
	}

	ht.CreatedResp(ctx, "get list customers", cWeb)
}
