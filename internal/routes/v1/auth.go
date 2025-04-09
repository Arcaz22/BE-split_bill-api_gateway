package v1

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
    {
		auth := router.Group("/auth")
		authHandler := handler.NewAuthHandler()

		auth.GET("/health", authHandler.Forward())
        auth.POST("/signup", authHandler.Forward())
        auth.POST("/signin", authHandler.Forward())
        auth.POST("/refresh-token", authHandler.Forward())
        auth.POST("/logout", authHandler.Forward())
        auth.GET("/verify", authHandler.Forward())
        auth.POST("/google", authHandler.Forward())
        auth.GET("/google/callback", authHandler.Forward())
    }

    user := router.Group("/user")
    user.Use(middleware.Auth())
    {
        authHandler := handler.NewAuthHandler()
        user.GET("/current-user", authHandler.Forward())
        user.GET("/get-all-user", authHandler.Forward())
        user.POST("/add-profile", authHandler.Forward())
        user.POST("/add-avatar-profile", authHandler.Forward())
        user.PUT("/update-profile", authHandler.Forward())
        user.PUT("/update-avatar-profile", authHandler.Forward())
    }

    role := router.Group("/role")
    role.Use(middleware.Auth())
    {
        authHandler := handler.NewAuthHandler()
        role.GET("/get-all-role", authHandler.Forward())
        role.POST("/create-role", authHandler.Forward())
        role.POST("/add-user-role", authHandler.Forward())
    }
}
