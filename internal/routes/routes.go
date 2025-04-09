package routes

import (
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes/v1"
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


    router.GET("/health", healthCheck)

	setupV1Routes(router)

	return router
}

func healthCheck(c *gin.Context) {
    response.JSON(c, http.StatusOK, gin.H{
        "status":    "OK",
        "timestamp": time.Now().Unix(),
        "message":   "API Gateway is running",
    })
}

func setupV1Routes(router *gin.Engine) {
    v1.SetupAuthRoutes(router)
    v1.SetupTransactionRoutes(router)
    v1.SetupNotificationRoutes(router)
}
