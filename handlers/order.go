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
// @Success 201 {object} web.NormalResp{data=web.PostOrderOutput}
// @Router /order [post]
func PostOrder(ctx *gin.Context) {
	var input web.PostOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NormalResp{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.NormalResp{
		Data: web.PostOrderOutput{},
	})
}
