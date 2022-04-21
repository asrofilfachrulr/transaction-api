package handlers

import "github.com/gin-gonic/gin"

// Register godoc
// @Summary add new user
// @Description  add new user
// @Tags User
// @Param Body body webModels.PostOrderInput true "Entry new order"
// @Produce json
// @Success 201 {object} webModels.WithDataResp{data=webModels.PostOrderOutput}
// @Router /order [post]
func PostOrder(ctx *gin.Context) {

}
