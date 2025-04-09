package v1

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(router *gin.Engine) {
	test := router.Group("/transaction")
	test.Use(middleware.Auth())
	{
		transactionHandler := handler.NewTransactionHandler()
		test.GET("/", transactionHandler.Forward())
		test.GET("/current-user", transactionHandler.Forward())
	}

	transaction := router.Group("/friendships")
	transaction.Use(middleware.Auth())
	{
		transactionHandler := handler.NewTransactionHandler()
		transaction.POST("/", transactionHandler.Forward())
		transaction.PUT("/:id", transactionHandler.Forward())
		transaction.GET("/:id", transactionHandler.Forward())
		transaction.GET("/", transactionHandler.Forward())
		transaction.DELETE("/:id", transactionHandler.Forward())
		transaction.GET("/pending", transactionHandler.Forward())
		transaction.GET("/friends", transactionHandler.Forward())
		transaction.PUT("/:id/accept", transactionHandler.Forward())
		transaction.PUT("/:id/reject", transactionHandler.Forward())
	}
}
