package handler

import (
    "api-gateway/config"
    "api-gateway/internal/proxy"
    "time"

    "github.com/gin-gonic/gin"
)

func NotificationHandler() gin.HandlerFunc {
    cfg := config.GetConfig()
    serviceURL := cfg.Services.Notification.URL
    timeout := cfg.Services.Notification.Timeout

    if timeout == 0 {
        timeout = 10 * time.Second
    }

    return proxy.ServiceForwarder(serviceURL, timeout)
}
