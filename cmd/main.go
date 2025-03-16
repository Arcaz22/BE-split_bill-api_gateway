package main

import (
	"log"
	"time"

	"api-gateway/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    if err := config.LoadConfig(); err != nil {
        log.Fatal("Cannot load config:", err)
    }

    router := gin.Default()

    middleware.InitCircuitBreaker()

    router.Use(middleware.Logger())
    router.Use(middleware.CORS())
    router.Use(middleware.RateLimit())
    router.Use(middleware.CircuitBreaker())

    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
			"status": "OK",
			"timestamp": time.Now().Unix(),
			"message": "API Gateway is running",
		})
    })

    routes.SetupRoutes(router)

    cfg := config.GetConfig()
    log.Fatal(router.Run(":" + cfg.Server.Port))
}
