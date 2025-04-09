package v1

import (
    "api-gateway/internal/handler"
    "api-gateway/internal/middleware"
    "github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(router *gin.Engine) {
	notifications := router.Group("/notifications")
    notifications.Use(middleware.Auth())
    {
        notificationHandler := handler.NewNotificationHandler()

		notifications.GET("/health", notificationHandler.HealthCheck)

        notifications.GET("/", notificationHandler.Forward())
        notifications.GET("/:notification_id", notificationHandler.Forward())
        notifications.PUT("/:notification_id/read", notificationHandler.Forward())
    }
}
