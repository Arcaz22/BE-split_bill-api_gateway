package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"api-gateway/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(middleware.RateLimit())
	router.Use(middleware.CircuitBreaker())

	router.GET("/health", func(c *gin.Context) {
		response.JSON(c, http.StatusOK, gin.H{
			"status":    "OK",
			"timestamp": time.Now().Unix(),
			"message":   "API Gateway is running",
		})
	})

	// Auth routes - public
	auth := router.Group("/auth")
	{
		authHandler := handler.AuthHandler()
		auth.POST("/signup", authHandler)
		auth.POST("/signin", authHandler)
		auth.POST("/refresh-token", authHandler)
		auth.POST("/logout", authHandler)
		auth.GET("/verify", authHandler)
		auth.POST("/google", authHandler)
		auth.GET("/google/callback", authHandler)
	}

	user := router.Group("/user")
	user.Use(middleware.Auth())
	{
		authHandler := handler.AuthHandler()
		user.GET("/current-user", authHandler)
		user.GET("/get-all-user", authHandler)
		user.POST("/add-profile", authHandler)
		user.POST("/add-avatar-profile", authHandler)
		user.PUT("/update-profile", authHandler)
		user.PUT("/update-avatar-profile", authHandler)
	}

	role := router.Group("/role")
	role.Use(middleware.Auth())
	{
		authHandler := handler.AuthHandler()
		role.GET("/get-all-role", authHandler)
		role.POST("/create-role", authHandler)
		role.POST("/add-user-role", authHandler)
	}

	// Transaction routes
	test := router.Group("/transaction")
	test.Use(middleware.Auth())
	{
		transactionHandler := handler.TransactionHandler()
		test.GET("/", transactionHandler)
		test.GET("/current-user", transactionHandler)
	}

	// Friendship routes
	transaction := router.Group("/friendships")
	transaction.Use(middleware.Auth())
	{
		transactionHandler := handler.TransactionHandler()
		transaction.POST("/", transactionHandler)
		transaction.PUT("/:id", transactionHandler)
		transaction.GET("/:id", transactionHandler)
		transaction.GET("/", transactionHandler)
		transaction.DELETE("/:id", transactionHandler)
		transaction.GET("/pending", transactionHandler)
		transaction.GET("/friends", transactionHandler)
		transaction.PUT("/:id/accept", transactionHandler)
		transaction.PUT("/:id/reject", transactionHandler)
	}

	// Notification routes
	notifications := router.Group("/notifications")
    notifications.Use(middleware.Auth())
    {
        notificationHandler := handler.NotificationHandler()
        notifications.GET("/", notificationHandler)
        notifications.GET("/:notification_id", notificationHandler)
        notifications.PUT("/:notification_id/read", notificationHandler)
    }

	return router
}
