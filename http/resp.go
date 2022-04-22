package http

import (
	"net/http"

	"github.com/asrofilfachrulr/transaction-api/models/web"
	"github.com/gin-gonic/gin"
)

func ErrResp(ctx *gin.Context, msg string, code ...int) {
	status := http.StatusBadRequest
	if len(code) == 0 {
		status = code[0]
	}

	ctx.JSON(status, web.NormalResp{
		Status:  "error",
		Message: msg,
	})
}

func OKResp(ctx *gin.Context, msg string, data ...any) {
	status := http.StatusOK

	ctx.JSON(status, web.NormalResp{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}

func CreatedResp(ctx *gin.Context, msg string, data ...any) {
	status := http.StatusCreated

	ctx.JSON(status, web.NormalResp{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}
