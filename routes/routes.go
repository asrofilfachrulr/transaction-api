package routes

import (
	"log"
	"net/http"

	"github.com/asrofilfachrulr/transaction-api/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Attach(r *gin.Engine, key string, data any) {
	r.Use(func(ctx *gin.Context) {
		ctx.Set(key, data)
	})
}

func InitAPI(r *gin.Engine) *gin.Engine {
	log.Println("Initiating route API...")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := r.Group("/api/v1")

	api.POST("/customer", handlers.PostCustomer)
	api.GET("/customers", handlers.GetAllCustomers)
	api.GET("/products", handlers.GetAllProducts)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
