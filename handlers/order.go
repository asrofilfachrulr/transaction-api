package handlers

import (
	"net/http"

	"github.com/asrofilfachrulr/transaction-api/models/web"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary add new user
// @Description  add new user
// @Tags User
// @Param Body body web.PostOrderInput true "Entry new order"
// @Produce json
// @Success 201 {object} web.WithDataResp{data=web.PostOrderOutput}
// @Fail 400 {object} web.PlainErr
// @Router /order [post]
func PostOrder(ctx *gin.Context) {
	var input web.PostOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.PlainErr{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.WithDataResp{
		Data: web.PostOrderOutput{},
	})
}
