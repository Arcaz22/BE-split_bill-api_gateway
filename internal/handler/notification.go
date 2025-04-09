package handler

import (
	"api-gateway/config"
	"api-gateway/internal/proxy"
	"api-gateway/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	serviceURL string
	timeout    time.Duration
}

func NewNotificationHandler() *NotificationHandler {
    cfg := config.GetConfig()
    timeout := cfg.Services.Notification.Timeout

    if timeout == 0 {
        timeout = 10 * time.Second
    }

    return &NotificationHandler{
        serviceURL: cfg.Services.Notification.URL,
        timeout:    timeout,
    }
}

// @Summary     Notification service health check
// @Description Check if notification service is running
// @Tags        system
// @Accept      json
// @Produce     json
// @Router      /notifications/health [get]
func (h *NotificationHandler) HealthCheck(c *gin.Context) {
	client := &http.Client{
		Timeout: h.timeout,
	}

	resp, err := client.Get(h.serviceURL + "/health")
	if err != nil {
		response.JSON(c, http.StatusServiceUnavailable, gin.H{
            "status":    "DOWN",
            "service":   "notification-service",
            "error":     err.Error(),
            "timestamp": time.Now().Format(time.RFC3339),
        })
        return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        response.JSON(c, http.StatusServiceUnavailable, gin.H{
            "status":    "DOWN",
            "service":   "notification-service",
            "code":      resp.StatusCode,
            "timestamp": time.Now().Format(time.RFC3339),
        })
        return
    }

    response.JSON(c, http.StatusOK, gin.H{
        "status":    "UP",
        "service":   "notification-service",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}

func (h *NotificationHandler) Forward() gin.HandlerFunc {
	return proxy.ServiceForwarder(h.serviceURL, h.timeout)
}
